package com.networkflow.backendspringboot3.model.request;

import lombok.Data;

import java.time.LocalDateTime;


@Data
public class TaskRequest {
    private String task_id;
    private LocalDateTime create_time;
    private LocalDateTime start_time;
    private LocalDateTime end_time;

    // mode: 0为实时流量检测，1为离线流量检测
    private Integer mode;
    // model: 0为XGBoost二分类模型, 1为XGBoost多分类模型, 2为Whisper二分类模型
    private  Integer model;
    // port: 实时流量检测端口
    private Integer port;
    // PCAP包存储位置
    private String pcap_path;
    // PCAP包存储位置(任务id命名)
    private String true_pcap_path;
    // 正常数据包数
    private Integer normal;
    // 异常数据包数
    private Integer abnormal;
    // 总数据包数
    private Integer total;

    // status：0 未启动；1 等待解析中；2 正在解析和检测中；3 完成解析聚合，正在检测中；4 正在汇总中；5 已完成。
    private Integer status;

}
