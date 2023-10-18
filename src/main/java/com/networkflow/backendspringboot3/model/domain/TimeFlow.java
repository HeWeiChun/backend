package com.networkflow.backendspringboot3.model.domain;

import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import lombok.Data;
import org.apache.commons.lang3.builder.ToStringBuilder;

import java.time.LocalDateTime;

@TableName(value = "time_flow")
@Data
public class TimeFlow {
    @TableId
    private String flow_id;
    private Integer ran_ue_ngap_id;
    private Integer total_num;
    private LocalDateTime begin_time;
    private LocalDateTime latest_time;
    private Integer verification_tag;
    private String src_ip;
    private String dst_ip;
    private Integer status_flow;
    private String task_id;
    @Override
    public String toString() {
        return ToStringBuilder.reflectionToString(this);
    }
}
