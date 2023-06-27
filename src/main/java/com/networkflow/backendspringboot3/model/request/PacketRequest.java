package com.networkflow.backendspringboot3.model.request;

import lombok.Data;

import java.time.LocalDateTime;
@Data
public class PacketRequest {
    private String NgapType;
    private String NgapProcedureCode;
    private Integer RanUeNgapId;
    private Integer PacketLen;
    private Integer ArriveTimeUs;
    private LocalDateTime ArriveTime;
    private Integer TimeInterval;
    private Integer VerificationTag;
    private String SrcIP;
    private String DstIP;
    private Integer DirSeq;
    private String FlowUEID;
    private String FlowTimeID;
    private Integer StatusPacket;
}
