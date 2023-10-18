package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.PacketMapper;
import com.networkflow.backendspringboot3.mapper.TaskMapper;
import com.networkflow.backendspringboot3.model.domain.Packet;
import com.networkflow.backendspringboot3.service.PacketService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.math.BigInteger;

@Service
public class PacketServiceImpl extends ServiceImpl<PacketMapper, Packet> implements PacketService {
    @Autowired
    private PacketMapper packetMapper;


    @Override
    public R allPacket() {
        QueryWrapper<Packet> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByAsc(Packet::getArrive_time);
        return R.success(null, packetMapper.selectList(queryWrapper));
    }

    @Override
    public R getPacketByFlowId(String flowId, Integer model) {
        QueryWrapper<Packet> queryWrapper = new QueryWrapper<>();
        if(model == 1)
            queryWrapper.lambda().eq(Packet::getFlow_time_id, flowId).orderByAsc(Packet::getArrive_time);
        else
            queryWrapper.lambda().eq(Packet::getFlow_ue_id, flowId).orderByAsc(Packet::getArrive_time);
        return R.success(null, packetMapper.selectList(queryWrapper));
    }
}
