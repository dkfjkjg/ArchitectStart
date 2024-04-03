package org.example.servicehi;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class ServiceHiApplication {

    @RequestMapping(value = "/hi")
    public String hello(String name) {
        return "hello " + name + "!";
    }

    public static void main(String[] args) {
        SpringApplication.run(ServiceHiApplication.class, args);
    }

}
