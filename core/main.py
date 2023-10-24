# -*- coding: UTF-8 -*-
import argparse
import subprocess
import os
import datetime
from clickhouse_driver import Client
import sys


client = Client(host='10.3.242.84', port=9000, user='default', password='password')


def main(args):
    update_status = "ALTER TABLE sctp.task UPDATE status = %(status)s WHERE task_id = %(taskid)s"
    taskid_param = {'taskid': args.taskid}
    try:
        # 执行Go脚本
        select_pcap = "SELECT true_pcap_path FROM sctp.task WHERE task_id = %(taskid)s"
        pcapPath = client.execute(select_pcap, taskid_param)[0][0]
        go_cmd = [
            "go",
            "run",
            "main.go",
            "--pcap_path",
            f"../upload/{pcapPath}",
            "--taskid",
            args.taskid
        ]
        go_process = subprocess.Popen(go_cmd, cwd="/home/whe/webHWCYahong/backend/core/go",
                                      stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        print(f"Go脚本进程PID为: {go_process.pid}")

        while True:
            line = go_process.stdout.readline()
            if not line:
                break
            # print("脚本输出: (Go) " + line.strip().decode('utf-8'))

        go_exit_code = go_process.wait()

        print(f"Go脚本执行完毕, 退出码: {go_exit_code}")

        if go_exit_code == 0:
            task_status = 3
        else:
            task_status = 100

        go_update_params = {'taskid': args.taskid, 'status': task_status}
        client.execute(update_status, go_update_params)
        if go_exit_code == 0:
            print("解析完成")
        else:
            print("解析失败")
            return

        # 执行Python脚本
        python_cmd = [
            "/home/whe/anaconda3/envs/xgboost39/bin/python",
            os.path.join(os.getcwd(), "core", "python", "springboot.py"),
            "--taskid",
            args.taskid,
            "--model",
            args.model
        ]
        python_process = subprocess.Popen(python_cmd, cwd="/home/whe/webHWCYahong/backend",
                                          stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
        print(f"Python脚本进程的PID为: {python_process.pid}")

        while True:
            line = python_process.stdout.readline()
            if not line:
                break
            print("(Python) " + line.strip().decode('utf-8'))

        python_exit_code = python_process.wait()
        print(f"Python脚本执行完毕, 退出码：{python_exit_code}")

        current_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        taskid_param = {'taskid': args.taskid}
        if args.model == '1':
            get_abnormal_flow = "SELECT COUNT(*) FROM sctp.time_flow WHERE task_id = %(taskid)s AND status_flow = 200"
            get_normal_flow = "SELECT COUNT(*) FROM sctp.time_flow WHERE task_id = %(taskid)s AND status_flow = 100"
        else:
            get_abnormal_flow = "SELECT COUNT(*) FROM sctp.ue_flow WHERE task_id = %(taskid)s AND status_flow = 200"
            get_normal_flow = "SELECT COUNT(*) FROM sctp.ue_flow WHERE task_id = %(taskid)s AND status_flow = 100"
        update_info = "ALTER TABLE sctp.task UPDATE status = %(status)s, normal = %(normal)s, " \
                      "abnormal = %(abnormal)s, total = %(total)s, end_time = %(endtime)s WHERE task_id = %(taskid)s"

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
            print("检测完成")
        else:
            print("检测失败")

    except Exception as e:
        print("脚本输出: " + e)
        task_status = 100
        fail_params = {'taskid': args.taskid, 'status': task_status}
        client.execute(update_status, fail_params)
        print("脚本检测失败")
        sys.exit()


if __name__ == '__main__':
    sys.stdout.reconfigure(encoding='utf-8')
    sys.stderr.reconfigure(encoding='utf-8')
    parser = argparse.ArgumentParser()
    parser.add_argument('--model', required=True, type=str)
    parser.add_argument('--taskid', required=True, type=str)
    args = parser.parse_args()
    main(args)
