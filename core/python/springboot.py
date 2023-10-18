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
import sys
sys.stdout.reconfigure(encoding='utf-8')
client = Client(host='10.3.242.84', port=9000, user='default', password='password')


def detect_taskid(model_type, taskid):
    # 加载模型
    loss_p = 60000

    if model_type == '0':  # XGBoost(UEID聚合)
        model_path = "core/Python/model/ueid_XGBoost1010.pkl"
        print("当前使用模型: XGBoost(UEID)")
    elif model_type == '1':  # XGBoost(Time聚合)
        model_path = "core/Python/model/time_XGBoost1010.pkl"
        print("当前使用模型: XGBoost(Time)")
    elif model_type == '2':  # Whipser模型(UEID聚合)
        model_path = "core/Python/model/kmeans.pkl"
        print("当前使用模型: Whipser(UEID)")
    else:
        model_path = "core/Python/model/ueid_XGBoost1010.pkl"
        print("不合法的模型类型. 使用默认模型: XGBoost(UEID)")
    model = joblib.load(model_path)
    random.seed(27)
    taskid_params = {'taskid': taskid}
    select = "SELECT status FROM sctp.task WHERE task_id = %(taskid)s"

    while True:
        # 获取当前任务状态: 0-未开始 1-待解析 2-解析中 3-待检测 4-检测中 5-检测完成
        task_status = client.execute(select, taskid_params)[0][0]
        if task_status != 3:
            time.sleep(1)
        elif task_status == 3:
            # 更新任务状态为检测中
            tasking = "ALTER TABLE sctp.task UPDATE status = 4 WHERE task_id = %(taskid)s"
            client.execute(tasking, taskid_params)
            # 获取当前任务的所有未检测的流ID
            if model_type == '1':
                flow = "SELECT flow_id FROM sctp.time_flow WHERE task_id = %(taskid)s AND status_flow = 0"
            else:
                flow = "SELECT flow_id FROM sctp.ue_flow WHERE task_id = %(taskid)s AND status_flow = 0"
            flow_id = client.execute(flow, taskid_params)

        if model_type == '2':
            for id in flow_id:
                # 获取当前流的所有包
                packet = ("SELECT packet_len,time_interval,ngap_type FROM sctp.packet "
                          "WHERE flow_ue_id = %(flowid)s ORDER BY arrive_time")
                packet_params = {'flowid': id}
                result = client.execute(packet, packet_params)
                # 特征提取 & 模型检测
                feature = whisper.extraction(result)
                feature = np.array(feature)
                f = open("core/Python/model/train_loss.data", 'rb')
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
                update_flow_query = 'ALTER TABLE sctp.ue_flow UPDATE status_flow = %(ypredict)s WHERE flow_id = %(flowid)s'
                result_params = {
                    'ypredict': predict_code,
                    'flowid': id
                }
                client.execute(update_flow_query, result_params)
            break
        else:
            if model_type == '1':
                packet = "SELECT * FROM sctp.packet WHERE flow_time_id = %(flowid)s ORDER BY arrive_time"
                update_flow_query = 'ALTER TABLE sctp.time_flow UPDATE status_flow = %(ypredict)s WHERE flow_id = %(flowid)s'
            else:
                packet = "SELECT * FROM sctp.packet WHERE flow_ue_id = %(flowid)s ORDER BY arrive_time"
                update_flow_query = 'ALTER TABLE sctp.ue_flow UPDATE status_flow = %(ypredict)s WHERE flow_id = %(flowid)s'

            for id in flow_id:
                # 获取当前流的所有包
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
