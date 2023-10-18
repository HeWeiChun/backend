package com.networkflow.backendspringboot3.model.domain;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;

import java.time.LocalDateTime;

@TableName(value = "packet")
@Data
public class Packet {
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

    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }
}
