package com.networkflow.backendspringboot3.service.impl;

import cn.hutool.log.Log;
import cn.hutool.log.LogFactory;
import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.networkflow.backendspringboot3.common.R;
import com.networkflow.backendspringboot3.mapper.PacketMapper;
import com.networkflow.backendspringboot3.mapper.TaskMapper;
import com.networkflow.backendspringboot3.mapper.TimeFlowMapper;
import com.networkflow.backendspringboot3.mapper.UEFlowMapper;
import com.networkflow.backendspringboot3.model.domain.Packet;
import com.networkflow.backendspringboot3.model.domain.Task;
import com.networkflow.backendspringboot3.model.domain.TimeFlow;
import com.networkflow.backendspringboot3.model.domain.UEFlow;
import com.networkflow.backendspringboot3.model.request.TaskRequest;
import com.networkflow.backendspringboot3.service.TaskService;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStreamReader;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.Executor;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

@Service
public class TaskServiceImpl extends ServiceImpl<TaskMapper, Task> implements TaskService {
    private static final Log log = LogFactory.get();
    private final DetectTask detectTask;
    @Autowired
    private TaskMapper taskMapper;
    @Autowired
    private TimeFlowMapper timeFlowMapper;
    @Autowired
    private UEFlowMapper ueFlowMapper;
    @Autowired
    private PacketMapper packetMapper;

    public TaskServiceImpl(DetectTask detectTask) {
        this.detectTask = detectTask;
    }

    @Override
    public R allTask() {
        QueryWrapper<Task> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().orderByDesc(Task::getCreateTime);
        return R.success(null, taskMapper.selectList(queryWrapper));
    }

    // 上传文件(返回任务id命名的名字)
    private String uploadFile(MultipartFile uploadFile, String taskId) {
        if (uploadFile == null) {
            return null;
        }
        String fileName = uploadFile.getOriginalFilename();
        // 求文件后缀
        String extension = "";
        if (fileName != null) {
            int dotIndex = fileName.lastIndexOf('.');
            if (dotIndex != -1 && dotIndex < fileName.length() - 1) {
                extension = fileName.substring(dotIndex + 1);
            }
        }
        String trueFileName = taskId + "." + extension;

        // 检查文件存储位置是否存在
        String filePath = System.getProperty("user.dir") + System.getProperty("file.separator") + "core_go" + System.getProperty("file.separator") + "upload";
        File file = new File(filePath);
        if (!file.exists()) {
            if (!file.mkdir()) {
                return null;
            }
        }
        // 文件路径
        File dest = new File(filePath + System.getProperty("file.separator") + trueFileName);
        try {
            uploadFile.transferTo(dest);
            return trueFileName;
        } catch (IOException e) {
            return null;
        }
    }

    private void deleteCache(String taskId) {
        timeFlowMapper.delete(new QueryWrapper<TimeFlow>().lambda().eq(TimeFlow::getTaskID, taskId));
        List<UEFlow> ueFlowList = ueFlowMapper.selectList(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getTaskID, taskId));
        String[] flowIds = ueFlowList.stream()
                .map(UEFlow::getFlowId)
                .toArray(String[]::new);
        if(flowIds.length>0)
            packetMapper.delete(new QueryWrapper<Packet>().lambda().in(Packet::getFlowUEID, Arrays.asList(flowIds)));
        ueFlowMapper.delete(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getTaskID, taskId));
    }

    @Override
    public R createTask(TaskRequest createTaskRequest, MultipartFile uploadFile) {
        Task task = new Task();

        BeanUtils.copyProperties(createTaskRequest, task);

        String trueFileName = uploadFile(uploadFile, createTaskRequest.getTaskId());
        if (trueFileName != null) {
            task.setPcapPath(uploadFile.getOriginalFilename());
            task.setTruePcapPath(trueFileName);
        } else {
            return R.fatal("上传文件失败");
        }

        if (taskMapper.insert(task) > 0) {
            return R.success("添加成功");
        } else {
            return R.error("添加失败");
        }
    }

    @Override
    public R updateTask(TaskRequest createTaskRequest, MultipartFile uploadFile) {
        Task task = new Task();
        BeanUtils.copyProperties(createTaskRequest, task);

        String trueFileName = uploadFile(uploadFile, createTaskRequest.getTaskId());
        if (trueFileName != null) {
            task.setPcapPath(uploadFile.getOriginalFilename());
            task.setTruePcapPath(trueFileName);
        } else {
            return R.fatal("上传文件失败");
        }

        if (taskMapper.updateById(task) > 0) {
            return R.success("更新成功");
        } else {
            return R.error("更新失败");
        }
    }

    @Override
    public R updateTaskStatus(String taskId, Integer status) {
        Task task = new Task();
        task.setTaskId(taskId);
        task.setStatus(status);
        if (taskMapper.updateById(task) > 0) {
            return R.success("更新成功");
        } else {
            return R.error("更新失败");
        }
    }


    @Override
    public R deleteTask(String[] taskIds) {
        for (String taskId : taskIds) {
            deleteCache(taskId);
            // 检查文件存储位置是否存在
            String filePath = System.getProperty("user.dir") + System.getProperty("file.separator") + "core_go" +
                    System.getProperty("file.separator") + "upload" + System.getProperty("file.separator") + taskId + ".pcapng";
            File file = new File(filePath);
            if (file.exists()) {
                boolean deleted = file.delete();
                if (deleted) {
                    System.out.println("文件删除成功");
                } else {
                    System.out.println("文件删除失败");
                }
            } else {
                System.out.println("文件不存在");
            }
        }
        if (taskMapper.deleteBatchIds(Arrays.asList(taskIds)) > 0) {
            return R.success("删除成功");
        } else {
            return R.error("删除失败");
        }
    }

    @Override
    public boolean updateTaskByTask(Task task) {
        return taskMapper.updateById(task) > 0;
    }

    @Override
    public R startTask(String[] taskIds) {
        int successCount = 0;
        for (String taskId : taskIds) {
            Task task = new Task();
            task.setTaskId(taskId);
            task.setStatus(1);
            task.setStartTime(null);
            task.setEndTime(null);
            task.setAbnormal(null);
            task.setNormal(null);
            task.setTotal(null);
            // 清除缓存
            deleteCache(taskId);

            if (taskMapper.updateById(task) > 0) {
                successCount++;
            }
        }
        if (successCount == taskIds.length) {
            return R.success("开始成功");
        } else if (successCount > 0 && successCount < taskIds.length) {
            return R.success("部分开始成功");
        } else {
            return R.error("开始失败");
        }
    }

    @Scheduled(cron = "0/5 * *  * * ? ")
    @Override
    public void checkStatus() {
        log.info("轮询数据库, 线程名字为 = " + Thread.currentThread().getName());

        QueryWrapper<Task> queryWrapper = new QueryWrapper<>();
        queryWrapper.lambda().eq(Task::getStatus, 1);
        List<Task> list = taskMapper.selectList(queryWrapper);

        for (Task task : list) {
            String currentTime = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
            task.setStatus(2);
            task.setStartTime(LocalDateTime.parse(currentTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")));
            taskMapper.updateById(task);
            if (taskMapper.updateById(task) > 0) {
                detectTask.executeGoScript(task);
            } else {
                log.info("启动成功");
            }
        }
    }
}

@Component
class DetectTask {
    private static final Log log = LogFactory.get();
    @Autowired
    private TaskMapper taskMapper;
    @Autowired
    private UEFlowMapper ueFlowMapper;
    @Autowired
    private Executor checkTaskPool;

    @Async("checkTaskPool")
    public void executePythonScript(String scriptPath, Task currentTask) {
        log.info("执行Python, 线程名字为 = " + Thread.currentThread().getName());
        try {
            ProcessBuilder processBuilder = new ProcessBuilder("C:\\Users\\HorizonHe\\.conda\\envs\\xgboost39\\python.exe", scriptPath, "--taskid", currentTask.getTaskId(), "--model", String.valueOf(currentTask.getModel()));
            processBuilder.redirectErrorStream(true); // 合并标准输出和标准错误流
            Process process = processBuilder.start();

            // 处理脚本的输出
            BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
            String line;
            while ((line = reader.readLine()) != null) {
                log.info(line);
            }
            int exitCode = process.waitFor();
            log.info("Python脚本执行完毕, 退出码：" + exitCode);

            String currentTime = LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
            Task task = new Task();
            task.setTaskId(currentTask.getTaskId());
            if (exitCode == 0) {
                Long abnormalFlowAll = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 200).eq(UEFlow::getTaskID, currentTask.getTaskId()));
                Long normalFlowAll = ueFlowMapper.selectCount(new QueryWrapper<UEFlow>().lambda().eq(UEFlow::getStatusFlow, 100).eq(UEFlow::getTaskID, currentTask.getTaskId()));
                task.setStatus(5);
                task.setAbnormal(Math.toIntExact(abnormalFlowAll));
                task.setNormal(Math.toIntExact(normalFlowAll));
                task.setTotal(Math.toIntExact(abnormalFlowAll + normalFlowAll));
                task.setEndTime(LocalDateTime.parse(currentTime, DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss")));
            } else {
                task.setStatus(100);
                task.setAbnormal(null);
                task.setNormal(null);
                task.setTotal(null);
                task.setEndTime(null);
            }

            if (taskMapper.updateById(task) > 0) {
                if (exitCode == 0)
                    log.info("检测完成");
                else {
                    log.info("检测失败");
                }
            } else {
                log.info("检测失败");
            }
        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
            Task task = new Task();
            task.setTaskId(currentTask.getTaskId());
            task.setStatus(100);
            taskMapper.updateById(task);
            log.info("检测失败");
        } finally {
            if (checkTaskPool instanceof ThreadPoolExecutor) {
                ((ThreadPoolExecutor) checkTaskPool).remove(Thread.currentThread());
            }
        }
    }

    @Async("checkTaskPool")
    public void executeGoScript(Task currentTask) {
        log.info("执行Go, 线程名字为 = " + Thread.currentThread().getName());
        try {
            ProcessBuilder processBuilder = new ProcessBuilder("C:\\Users\\HorizonHe\\sdk\\go1.20.4\\bin\\go.exe", "run", "main.go", "--pcap_path", "..\\upload\\" + currentTask.getTruePcapPath(), "--taskid", currentTask.getTaskId());
            processBuilder.directory(new File("E:\\Code\\web\\backendspringboot3\\core_go\\sctp_flowmap"));
            processBuilder.redirectErrorStream(true); // 合并标准输出和标准错误流
            Process process = processBuilder.start();

            BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
            String line;
            while ((line = reader.readLine()) != null) {
                log.info(line);
            }

            int exitCode = process.waitFor();
            log.info("Go脚本执行完毕，退出码：" + exitCode);

            Task task = new Task();
            task.setTaskId(currentTask.getTaskId());
            if (exitCode == 0)
                task.setStatus(3);
            else {
                task.setStatus(100);
                task.setAbnormal(null);
                task.setNormal(null);
                task.setTotal(null);
                task.setEndTime(null);
            }

            if (taskMapper.updateById(task) > 0) {
                if (exitCode == 0)
                    log.info("解析完成");
                else {
                    log.info("解析失败");
                    return;
                }
            } else {
                log.info("解析失败");
                return;
            }

            executePythonScript(System.getProperty("user.dir") + System.getProperty("file.separator") + "core_go\\python" +
                    System.getProperty("file.separator") + "springboot.py", currentTask);

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();
            Task task = new Task();
            task.setTaskId(currentTask.getTaskId());
            task.setStatus(100);
            taskMapper.updateById(task);
            log.info("检测失败");
        }
    }
}