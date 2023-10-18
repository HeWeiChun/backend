package com.networkflow.backendspringboot3.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.model.domain.Packet;

public interface PacketService extends IService<Packet> {
    R allPacket();

    R getPacketByFlowId(String flowId, Integer model);
}
