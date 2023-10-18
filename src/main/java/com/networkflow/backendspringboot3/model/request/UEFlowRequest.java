package com.networkflow.backendspringboot3.model.request;

import lombok.Data;

import java.math.BigInteger;
import java.time.LocalDateTime;

@Data
public class UEFlowRequest {
    private String flow_id;
    private Integer ran_ue_ngap_id;
    private Integer total_num;
    private Integer start_second;
    private Integer end_second;
    private LocalDateTime begin_time;
    private LocalDateTime latest_time;
    private Integer verification_tag;
    private String src_ip;
    private String dst_ip;
    private Integer status_flow;
    private String task_id;
}
