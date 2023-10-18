package com.networkflow.backendspringboot3.model.request;

import lombok.Data;

import java.time.LocalDateTime;

@Data
public class TimeFlowRequest {
    private String flowId;
    private Integer ran_ue_ngap_id;
    private Integer total_num;
    private LocalDateTime begin_time;
    private LocalDateTime latest_time;
    private Integer verification_tag;
    private String src_ip;
    private String dst_ip;
    private Integer status_flow;
    private String task_id;
}
