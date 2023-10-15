package com.networkflow.backendspringboot3.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.*;
import com.networkflow.backendspringboot3.model.domain.Abstract;
import com.networkflow.backendspringboot3.model.domain.Task;
import com.networkflow.backendspringboot3.model.domain.TimeFlow;
import com.networkflow.backendspringboot3.model.domain.UEFlow;
import com.networkflow.backendspringboot3.service.AbstractService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.ArrayList;
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
        // task_status
        // status: 0(未开始),1(待解析),2(解析中),3(待检测),4(检测中),5(检测完成),100(错误)
        // flow_status
        // status: 0(未检测),100(检测完成且为正常)，200(检测完成且为异常)

        // 介绍栏
        // 活跃任务
        Map<String, Integer> activeTask = new HashMap<>();
        // 活跃任务——在线任务数(统计任务表中有多少mode为1的任务)
        Long online = taskMapper.selectCount(new QueryWrapper<Task>().lambda().eq(Task::getMode, 1).in(Task::getStatus,0,1,2,3,4));
        // 活跃任务——离线任务数(统计任务表中有多少mode为1的任务)
        Long offline = taskMapper.selectCount(new QueryWrapper<Task>().lambda().eq(Task::getMode, 0).in(Task::getStatus,0,1,2,3,4));
        activeTask.put("online", online.intValue());
        activeTask.put("offline", offline.intValue());

        // 已完成任务数(按每天计算)(数据库中endtime的时间精确到分, 以天为单位，返回每天进行了多少任务)
        Map<String, Integer> completedTask = new HashMap<>();
        // 异常流数(统计UEFlow和TimeFlow中共有多少status为200)
        Map<String, Long> n2Abnormal = new HashMap<>();
        // 正常流数(统计UEFlow和TimeFlow中共有多少status为100)
        Map<String, Long> n2Normal = new HashMap<>();
        List<Task> completedTasks = taskMapper.selectList(new QueryWrapper<Task>().lambda().eq(Task::getStatus, 5));
        for (Task task : completedTasks) {
            LocalDateTime createTime = task.getCreateTime();
            String day = createTime.toLocalDate().toString();
            String taskId = task.getTaskId();

            Long abnormalFlow = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 200).eq(UEFlow::getTaskID, taskId));
            Long normalFlow = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 100).eq(UEFlow::getTaskID, taskId));

            completedTask.put(day, completedTask.getOrDefault(day, 0) + 1);
            n2Abnormal.put(day, n2Abnormal.getOrDefault(day, 0L) + abnormalFlow);
            n2Normal.put(day, n2Normal.getOrDefault(day, 0L) + normalFlow);
        }
        // 转换completedTask
        List<Map<String, Object>> completedTaskList = new ArrayList<>();
        for (Map.Entry<String, Integer> entry : completedTask.entrySet()) {
            String dateByDay = entry.getKey();
            Integer count = entry.getValue();

            Map<String, Object> completedTaskItem = new HashMap<>();
            completedTaskItem.put("dateByDay", dateByDay);
            completedTaskItem.put("completeTaskByDay", count);

            completedTaskList.add(completedTaskItem);
        }

        // 转换n2Abnormal
        List<Map<String, Object>> n2AbnormalList = new ArrayList<>();
        for (Map.Entry<String, Long> entry : n2Abnormal.entrySet()) {
            String dateByDay = entry.getKey();
            Long n2AbnormalByDay = entry.getValue();

            Map<String, Object> n2AbnormalItem = new HashMap<>();
            n2AbnormalItem.put("dateByDay", dateByDay);
            n2AbnormalItem.put("n2AbnormalByDay", n2AbnormalByDay);

            n2AbnormalList.add(n2AbnormalItem);
        }

        // 转换n2Normal
        List<Map<String, Object>> n2NormalList = new ArrayList<>();
        for (Map.Entry<String, Long> entry : n2Normal.entrySet()) {
            String dateByDay = entry.getKey();
            Long n2NormalByDay = entry.getValue();

            Map<String, Object> n2NormalItem = new HashMap<>();
            n2NormalItem.put("dateByDay", dateByDay);
            n2NormalItem.put("n2NormalByDay", n2NormalByDay);

            n2NormalList.add(n2NormalItem);
        }


        // 所有流
        Long abnormalFlowAll = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 200));
        Long normalFlowAll = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 100));
        Map<String, Long> abnormalFlowBinary = new HashMap<>();
        abnormalFlowBinary.put("normal", normalFlowAll);
        abnormalFlowBinary.put("abnormal", abnormalFlowAll);

        Map<String, Long> abnormalFlowMulti = new HashMap<>();
        abnormalFlowMulti.put("normal", normalFlowAll);
        abnormalFlowMulti.put("abnormal", abnormalFlowAll);

        // 异常事件(返回UEFlow和TimeFlow中所有status为1的流，并以时间倒序排序)
        QueryWrapper<UEFlow> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(UEFlow::getStatusFlow, 200).orderByDesc(UEFlow::getLatestTime);
        List<UEFlow> abnormalEvent = ueFlowMapper.selectList(queryWrapper);


        // 活跃流数——已检测流(统计UEFlow和TimeFlow中共有多少status为0的流)
        Long activeDetectedFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().in("StatusFlow", 100, 200))
                + timeFlowMapper.selectCount(new QueryWrapper<TimeFlow>().in("StatusFlow", 100, 200));

        // 活跃流数——待检测流(统计UEFlow和TimeFlow中共有多少status为1的流)
        Long activePendingFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().eq("StatusFlow", 0))
                + timeFlowMapper.selectCount(new QueryWrapper<TimeFlow>().eq("StatusFlow", 0));


        Map<String, Object> introduce = new HashMap<>();
        introduce.put("activeTask", activeTask);
        introduce.put("completedTask", completedTaskList);
        introduce.put("n2Abnormal", n2AbnormalList);
        introduce.put("n2Normal", n2NormalList);

        Map<String, Object> result = new HashMap<>();
        result.put("introduce", introduce);
        result.put("abnormalFlowBinary", abnormalFlowBinary);
        result.put("abnormalFlowMulti", abnormalFlowMulti) ;
        result.put("abnormalEvent", abnormalEvent);

        return R.success("success", result);
    }
    @Override
    public R abstractByID(String taskId) {
        // task_status
        // status: 0(未开始),1(待解析),2(解析中),3(待检测),4(检测中),5(检测完成),100(错误)
        // flow_status
        // status: 0(未检测),100(检测完成且为正常)，200(检测完成且为异常)

        // 介绍栏
        // 所有流
        Long abnormalFlowAll = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 200).eq(UEFlow::getTaskID, taskId));
        Long normalFlowAll = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 100).eq(UEFlow::getTaskID, taskId));
        Map<String, Long> abnormalFlowBinary = new HashMap<>();
        abnormalFlowBinary.put("normal", normalFlowAll);
        abnormalFlowBinary.put("abnormal", abnormalFlowAll);

        Map<String, Long> abnormalFlowMulti = new HashMap<>();
        abnormalFlowMulti.put("normal", normalFlowAll);
        abnormalFlowMulti.put("abnormal", abnormalFlowAll);



        // 已检测流
        Long activeDetectedFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().in(UEFlow::getStatusFlow, 100, 200).eq(UEFlow::getTaskID, taskId));

        // 待检测流
        Long activePendingFlows = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 0).eq(UEFlow::getTaskID, taskId));



        Map<String, Object> result = new HashMap<>();
        result.put("activeDetectedFlows", activeDetectedFlows);
        result.put("activePendingFlows", activePendingFlows);
        result.put("abnormalFlowBinary", abnormalFlowBinary);
        result.put("abnormalFlowMulti", abnormalFlowMulti);

        return R.success("success", result);
    }
}