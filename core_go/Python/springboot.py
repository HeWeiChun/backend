import argparse
import pickle
import random
import time

import joblib
import pandas as pd
import numpy as np
from clickhouse_driver import Client

import module.feature_extraction as fe
import module.whisper_feature_extraction as whisper
import module.read_pcap as rp

client = Client(host='10.3.242.84', port=9000, user='default', password='password')


def detect(model_type, file_path, taskid):
    file_path = "E:\\Code\\web\\backendspringboot3\\core\\upload\\" + file_path
    suffix = file_path.split(".")[-1]
    if suffix == "pcap" or suffix == "pcapng":
        a = rp.read_pcap(file_path)
    elif suffix == "csv":
        a = pd.read_csv(file_path)
    X, y = fe.get_dataset(a, client, taskid)
    if model_type == 'bin':
        model_path = "/core/python/model/model_bin_UEID.pkl"
    else:
        model_path = "E:/Code/web/backendspringboot3/core_go/Python/model/model_multi_UEID.pkl"
    model = joblib.load(model_path)
    if len(X) > 0:
        y_predict = model.predict(X)
        print("Total traffic:", len(X), " Abnormal traffic: ", sum(y_predict >= 1))
        query = 'ALTER TABLE SCTP.Task UPDATE normal = %(normal)s, abnormal = %(abnormal)s, total = %(total)s WHERE taskId = %(taskid)s'
        params = {
            'normal': sum(y_predict == 0),
            'abnormal': sum(y_predict >= 1),
            'total': len(X),
            'taskid': taskid
        }
        client.execute(query, params)
    else:
        print("All normal traffic")
        query = 'ALTER TABLE SCTP.Task UPDATE normal = %(normal)s, abnormal = %(abnormal)s, total = %(total)s WHERE taskId = %(taskid)s'
        params = {
            'normal': 0,
            'abnormal': 0,
            'total': 0,
            'taskid': taskid
        }
        client.execute(query, params)
    client.disconnect()


def detect_taskid(model_type, taskid):
    # 加载模型
    loss_p = 10
    print(model_type)
    if model_type == '0':  # XGBoost二分类模型
        model_path = "core_go/Python/model/model_bin_UEID.pkl"
    elif model_type == '1':  # XGBoost多分类模型
        model_path = "core_go/Python/model/model_multi_UEID.pkl"
    elif model_type == '2':  # 多分类模型
        model_path = "core_go/Python/model/kmeans.pkl"
    else:
        model_path = "core_go/Python/model/model_bin_UEID.pkl"
    model = joblib.load(model_path)

    random.seed(27)
    taskid_params = {'taskid': taskid}
    select = "SELECT status FROM SCTP.Task WHERE taskId = %(taskid)s"

    while True:
        # 获取当前任务状态: 0-未开始 1-待解析 2-解析中 3-待检测 4-检测中 5-检测完成
        task_status = client.execute(select, taskid_params)[0][0]
        if task_status != 3:
            time.sleep(1)
        elif task_status == 3:
            # 更新任务状态为检测中
            tasking = "ALTER TABLE SCTP.Task UPDATE status = 4 WHERE taskId = %(taskid)s"
            client.execute(tasking, taskid_params)

            # 获取当前任务的所有未检测的流ID
            flow = "SELECT FlowId FROM SCTP.UEFlow WHERE TaskID = %(taskid)s AND StatusFlow = 0"
            flow_id = client.execute(flow, taskid_params)

        if model_type == '2':
            for id in flow_id:
                # 获取当前流的所有包
                packet = "SELECT PacketLen,TimeInterval,NgapType FROM SCTP.Packet WHERE FlowUEID = %(flowid)s ORDER BY ArriveTime"
                packet_params = {'flowid': id}
                result = client.execute(packet, packet_params)
                # 特征提取 & 模型检测
                feature = whisper.extraction(result)
                feature = np.array(feature)
                f = open("core_go/Python/model/train_loss.data", 'rb')
                train_loss = pickle.load(f)
                centers = model.cluster_centers_
                labels = model.predict(feature)
                prediction = []
                # print("开始测试数据...")
                print(train_loss)
                predict_code = 100
                for i in range(len(feature)):
                    temp = feature[i] - centers[labels[i]]
                    if np.linalg.norm(temp) > train_loss * loss_p:
                        # ANORMAL
                        prediction.append(0)
                        predict_code = 200
                    else:
                        # NORMAL
                        prediction.append(1)
                # 更新检测结果
                update_flow_query = 'ALTER TABLE SCTP.UEFlow UPDATE StatusFlow = %(ypredict)s WHERE FlowId = %(flowid)s'
                result_params = {
                    'ypredict': predict_code,
                    'flowid': id
                }
                client.execute(update_flow_query, result_params)
            break
        else:
            for id in flow_id:
                # 获取当前流的所有包
                packet = "SELECT * FROM SCTP.Packet WHERE FlowUEID = %(flowid)s ORDER BY ArriveTime"
                packet_params = {'flowid': id}
                result = client.execute(packet, packet_params)

                # 提取原始特征(包长, 时间戳, 序列方向)
                X = []
                for row in result:
                    dirseq = row[10]
                    if dirseq == 1:
                        dirseq = 1
                    else:
                        dirseq = -1
                    X.append([row[2], row[3], row[5], dirseq, 1])
                df = pd.DataFrame(X, columns=["RAN-UE-NGAP-ID", "Length", "Time", "DirSeq", "Label"])
                df['Time'] = df['Time'].astype('int64') / 10 ** 9

                # 特征提取 & 模型检测
                feature, label = fe.feature_extract(df)
                y_predict = model.predict([feature])[0]
                if y_predict == 0:
                    predict_code = 100
                else:
                    predict_code = 200
                # 更新检测结果
                update_flow_query = 'ALTER TABLE SCTP.UEFlow UPDATE StatusFlow = %(ypredict)s WHERE FlowId = %(flowid)s'
                result_params = {
                    'ypredict': predict_code,
                    'flowid': id
                }
                client.execute(update_flow_query, result_params)
            break

    client.disconnect()


def main(parser):
    args = parser.parse_args()
    # detect(args.model_type, args.file_path, args.taskid)
    detect_taskid(args.model, args.taskid)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--file_path', required=True, type=str)
    # parser.add_argument('--file_path',default=".\dataset\mix_dataset_test.csv",type=str)
    parser.add_argument('--model', required=True, type=str)
    parser.add_argument('--taskid', required=True, type=str)
    main(parser)
