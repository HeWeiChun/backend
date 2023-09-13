package com.networkflow.backendspringboot3;



import com.networkflow.backendspringboot3.model.domain.Task;
import com.networkflow.backendspringboot3.model.request.AdminRequest;
import com.networkflow.backendspringboot3.service.AbstractService;
import com.networkflow.backendspringboot3.service.impl.AbstractServiceImpl;
import com.networkflow.backendspringboot3.service.impl.AdminServiceImpl;
import com.networkflow.backendspringboot3.service.impl.TaskServiceImpl;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.scheduling.annotation.Scheduled;

import cn.hutool.log.LogFactory;
import cn.hutool.log.Log;

import java.io.BufferedReader;
import java.io.File;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.concurrent.TimeUnit;

@SpringBootTest
class Backendspringboot3ApplicationTests {

    @Autowired
    private TaskServiceImpl taskService;
    @Autowired
    private AbstractServiceImpl abstractService;
    private static final Log log = LogFactory.get();
    @Test
    void contextLoads() {
        System.out.println(taskService.allTask());
    }
    @Test
    void test1() {
        System.out.println(abstractService.allAbstract());
    }
    @Test
    void test2() {
        log.info("执行Go, 线程名字为 = " + Thread.currentThread().getName());
        try {
            ProcessBuilder processBuilder = new ProcessBuilder("C:\\Users\\HorizonHe\\sdk\\go1.20.4\\bin\\go.exe", "run", "main.go", "--pcap_path", "..\\upload\\" + "d383a92b-b6ea-4de1-9827-25214f1911dd.pcapng", "--taskid", "1111");
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

        } catch (IOException | InterruptedException e) {
            e.printStackTrace();

        }
    }

}
