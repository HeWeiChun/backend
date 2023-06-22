package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.*;
import com.networkflow.backendspringboot3.model.domain.Abstract;
import com.networkflow.backendspringboot3.model.domain.TimeFlow;
import com.networkflow.backendspringboot3.model.domain.UEFlow;
import com.networkflow.backendspringboot3.model.domain.Task;
import com.networkflow.backendspringboot3.service.AbstractService;
import org.apache.ibatis.javassist.bytecode.CodeIterator;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class AbstractServiceImpl extends ServiceImpl<AbstractMapper, Abstract> implements AbstractService {
    @Autowired
    private AbstractMapper abstractMapper;
    @Autowired
    private PacketMapper packetMapper;
    @Autowired
    private TaskMapper taskMapper;
    @Autowired
    private TimeFlowMapper timeFlowMapper;
    @Autowired
    private UEFlowMapper ueFlowMapper;

    @Override
    public R allAbstract() {
        List<DataPoint> offlineData = new ArrayList<>();
        List<DataPoint> visitData = new ArrayList<>();

        // task_status
        // status: 0(未开始),1(待解析),2(解析中),3(待检测),4(检测中),5(检测完成),100(错误)
        // flow_status
        // status: 0(未检测),100(检测完成且为正常)，200(检测完成且为异常)


        // 活跃任务——在线任务数(统计任务表中有多少mode为1的任务)
        Long activeOnlineTasks = taskMapper.selectCount(new QueryWrapper<Task>().eq("mode", 1));

        // 活跃任务——离线任务数(统计任务表中有多少mode为1的任务)
        Long activeOfflineTasks = taskMapper.selectCount(new QueryWrapper<Task>().eq("mode", 0));

        // 已完成任务数(按每天计算)(数据库中endtime的时间精确到分, 以天为单位，返回每天进行了多少任务)
        Map<String, Integer> completedTasksByDay = new HashMap<>();

        List<Task> completedTasks = taskMapper.selectList(new QueryWrapper<Task>().eq("status", 5));
        for (Task task : completedTasks) {
            LocalDateTime endTime = task.getEndTime();
            String day = endTime.toLocalDate().toString();
            completedTasksByDay.put(day, completedTasksByDay.getOrDefault(day, 0) + 1);
        }

        // 活跃流数——已检测流(统计UEFlow和TimeFlow中共有多少status为0的流)
        Long activeDetectedFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().in("StatusFlow", 100, 200))
                + timeFlowMapper.selectCount(new QueryWrapper<TimeFlow>().in("StatusFlow", 100, 200));

        // 活跃流数——待检测流(统计UEFlow和TimeFlow中共有多少status为1的流)
        Long activePendingFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().eq("StatusFlow", 0))
                + timeFlowMapper.selectCount(new QueryWrapper<TimeFlow>().eq("StatusFlow", 0));

        // 异常流数(统计UEFlow和TimeFlow中共有多少status为0且type=1的流)
        Long abnormalFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().eq("StatusFlow", 200))
                + timeFlowMapper.selectCount(new QueryWrapper<TimeFlow>().eq("StatusFlow", 200));

        // 正常流数(统计UEFlow和TimeFlow中共有多少status为0且type=2的流)
        Long normalFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().eq("StatusFlow", 100))
                + timeFlowMapper.selectCount(new QueryWrapper<TimeFlow>().eq("StatusFlow", 100));

//        // 异常事件(返回UEFlow和TimeFlow中所有status为1的流，并以时间倒序排序)
//        List<Event> abnormalEvents = new ArrayList<>();
//
//        List<UEFlow> ueFlowList = ueFlowMapper.selectList(new QueryWrapper<UEFlow>().eq("status", 1));
//        List<TimeFlow> timeFlowList = timeFlowMapper.selectList(new QueryWrapper<TimeFlow>().eq("status", 1));
//
//        for (UEFlow ueFlow : ueFlowList) {
//            Event event = new Event(ueFlow.getTimestamp(), ueFlow);
//            abnormalEvents.add(event);
//        }
//
//        for (TimeFlow timeFlow : timeFlowList) {
//            Event event = new Event(timeFlow.getTimestamp(), timeFlow);
//            abnormalEvents.add(event);
//        }
//
//        Collections.sort(abnormalEvents, Comparator.comparing(Event::getTimestamp).reversed());


        // 添加示例数据点
        visitData.add(new DataPoint("2023-06-16", 7));
        visitData.add(new DataPoint("2023-06-17", 8));
        visitData.add(new DataPoint("2023-06-18", 9));
        visitData.add(new DataPoint("2023-06-19", 10));
        visitData.add(new DataPoint("2023-06-20", 11));
        offlineData.add(new DataPoint("Stores 0", 0.2));
        offlineData.add(new DataPoint("Stores 1", 0.7));

        Map<String, Object> result = new HashMap<>();
        result.put("activeOnlineTasks", activeOnlineTasks);
        result.put("activeOfflineTasks", activeOfflineTasks);
        result.put("completedTasksByDay", completedTasksByDay);
        result.put("activeDetectedFlows", activeDetectedFlows);
        result.put("activePendingFlows", activePendingFlows);
        result.put("abnormalFlows", abnormalFlows);
        result.put("normalFlows", normalFlows);
//        result.put("abnormalEvents", abnormalEvents);
        result.put("visitData", visitData);
        result.put("offlineData", offlineData);
//        result.put("ttt", ueFlowList);

        return R.success("success", result);
    }

}
class DataPoint {
    @JsonProperty("x")
    private String x;
    @JsonProperty("y")
    private double y;

    public DataPoint(String x, double y) {
        this.x = x;
        this.y = y;
    }
}