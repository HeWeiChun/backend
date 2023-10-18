package com.networkflow.backendspringboot3.model.request;

import lombok.Data;

import java.time.LocalDateTime;
@Data
public class PacketRequest {
    private String ngap_type;
    private String ngap_procedure_code;
    private Integer ran_ue_ngap_id;
    private Integer packet_len;
    private Integer arrive_time_us;
    private LocalDateTime arrive_time;
    private Integer time_interval;
    private Integer verification_tag;
    private String src_ip;
    private String dst_ip;
    private Integer dir_seq;
    private String flow_ue_id;
    private String flow_time_id;
    private Integer status_packet;
}
