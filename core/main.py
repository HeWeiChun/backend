# -*- coding: UTF-8 -*-
import argparse
import subprocess
import os
import datetime
from clickhouse_driver import Client
import sys
sys.stdout.reconfigure(encoding='utf-8')

client = Client(host='10.3.242.84', port=9000, user='default', password='password')


def main(args):
    update_status = "ALTER TABLE SCTP.Task UPDATE status = %(status)s WHERE taskId = %(taskid)s"
    taskid_param = {'taskid': args.taskid}
    try:
        # 执行Go脚本
        select_pcap = "SELECT truePcapPath FROM SCTP.Task WHERE taskId = %(taskid)s"
        pcapPath = client.execute(select_pcap, taskid_param)[0][0]
        go_cmd = [
            "C:\\Users\\HorizonHe\\sdk\\go1.20.4\\bin\\go.exe",
            "run",
            "main.go",
            "--pcap_path",
            f"..\\upload\\{pcapPath}",
            "--taskid",
            args.taskid
        ]
        go_process = subprocess.Popen(go_cmd, cwd="E:\\Code\\web\\backendspringboot3\\core\\Go",
                                      stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        print(f"脚本输出: Go脚本进行的PID为: {go_process.pid}")

        while True:
            line = go_process.stdout.readline()
            if not line:
                break
            # print("脚本输出: (Go) " + line.strip().decode('utf-8'))

        go_exit_code = go_process.wait()

        print(f"脚本输出: Go脚本执行完毕，退出码：{go_exit_code}")

        if go_exit_code == 0:
            task_status = 3
        else:
            task_status = 100

        go_update_params = {'taskid': args.taskid, 'status': task_status}
        client.execute(update_status, go_update_params)
        if go_exit_code == 0:
            print("脚本输出: 解析完成")
        else:
            print("脚本输出: 解析失败")
            return

        # 执行Python脚本
        python_cmd = [
            "C:\\Users\\HorizonHe\\.conda\\envs\\xgboost39\\python.exe",
            os.path.join(os.getcwd(), "core", "python", "springboot.py"),
            "--taskid",
            args.taskid,
            "--model",
            args.model
        ]
        python_process = subprocess.Popen(python_cmd, cwd="E:\\Code\\web\\backendspringboot3",
                                          stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        print(f"脚本输出: Python脚本进行的PID为: {python_process.pid}")

        while True:
            line = python_process.stdout.readline()
            if not line:
                break
            print("脚本输出: (Python) " + line.strip().decode('utf-8'))

        python_exit_code = python_process.wait()
        print(f"脚本输出: Python脚本执行完毕, 退出码：{python_exit_code}")

        current_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        taskid_param = {'taskid': args.taskid}
        if args.model == '1':
            get_abnormal_flow = "SELECT COUNT(*) FROM SCTP.TimeFlow WHERE TaskID = %(taskid)s AND StatusFlow = 200"
            get_normal_flow = "SELECT COUNT(*) FROM SCTP.TimeFlow WHERE TaskID = %(taskid)s AND StatusFlow = 100"
        else:
            get_abnormal_flow = "SELECT COUNT(*) FROM SCTP.UEFlow WHERE TaskID = %(taskid)s AND StatusFlow = 200"
            get_normal_flow = "SELECT COUNT(*) FROM SCTP.UEFlow WHERE TaskID = %(taskid)s AND StatusFlow = 100"
        update_info = "ALTER TABLE SCTP.Task UPDATE status = %(status)s, normal = %(normal)s, " \
                      "abnormal = %(abnormal)s, total = %(total)s, endTime = %(endtime)s WHERE taskId = %(taskid)s"

        if python_exit_code == 0:
            abnormal_flow = client.execute(get_abnormal_flow, taskid_param)[0][0]
            normal_flow = client.execute(get_normal_flow, taskid_param)[0][0]
            task_status = 5
            update_info_params = {
                'status': task_status,
                'normal': normal_flow,
                'abnormal': abnormal_flow,
                'total': abnormal_flow + normal_flow,
                'endtime': current_time,
                'taskid': args.taskid
            }
            client.execute(update_info, update_info_params)
        else:
            task_status = 100
            update_info_params = {
                'status': task_status,
                'normal': None,
                'abnormal': None,
                'total': None,
                'endtime': None,
                'taskid': args.taskid
            }
            client.execute(update_info, update_info_params)

        if task_status == 5:
            print("脚本输出: 检测完成")
        else:
            print("脚本输出: 检测失败")

    except Exception as e:
        print("脚本输出: " + e)
        task_status = 100
        fail_params = {'taskid': args.taskid, 'status': task_status}
        client.execute(update_status, fail_params)
        print("脚本输出: 脚本检测失败")


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--model', required=True, type=str)
    parser.add_argument('--taskid', required=True, type=str)
    args = parser.parse_args()
    main(args)
