import argparse
import pickle
import random
import time

import joblib
import numpy as np
import pandas as pd
from clickhouse_driver import Client

import module.feature_extraction as fe
import module.whisper_feature_extraction as whisper

client = Client(host='10.3.242.84', port=9000, user='default', password='password')


def detect_taskid(model_type, taskid):
    # 加载模型
    loss_p = 60000

    if model_type == '0':  # XGBoost(UEID聚合)
        model_path = "core_go/Python/model/ueid_XGBoost1010.pkl"
        print("Current model: XGBoost(UEID)")
    elif model_type == '1':  # XGBoost(Time聚合)
        model_path = "core_go/Python/model/time_XGBoost1010.pkl"
        print("Current model: XGBoost(Time)")
    elif model_type == '2':  # Whipser模型(UEID聚合)
        model_path = "core_go/Python/model/kmeans.pkl"
        print("Current model: Whipser(UEID)")
    else:
        model_path = "core_go/Python/model/ueid_XGBoost1010.pkl"
        print("Invalid model type. Use default model: XGBoost(UEID)")
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
                packet = ("SELECT PacketLen,TimeInterval,NgapType FROM SCTP.Packet "
                          "WHERE FlowUEID = %(flowid)s ORDER BY ArriveTime")
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
                # print(train_loss)
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
                    X.append([row[1], row[2], row[3], row[5], row[10], row[13], row[14], row[15]])
                df = pd.DataFrame(X, columns=["ProcedureCode", "RAN-UE-NGAP-ID", "PacketLen", "Time", "DirSeq",
                                              "InitiatingMessage", "SuccessfulOutcome", "UnsuccessfulOutcome"])
                df['Time'] = df['Time'].astype('int64') / 10 ** 9

                # 特征提取 & 模型检测
                if len(df) > 1:
                    feature = fe.ngap_feature_extract(df)
                    y_predict = model.predict([feature])[0]
                else:
                    y_predict = 0
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
    detect_taskid(args.model, args.taskid)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--model', required=True, type=str)
    parser.add_argument('--taskid', required=True, type=str)
    main(parser)
