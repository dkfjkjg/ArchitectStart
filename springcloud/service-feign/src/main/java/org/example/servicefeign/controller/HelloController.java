package org.example.servicefeign.controller;

import org.example.servicefeign.service.ServiceHi;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
public class HelloController {

    @Autowired
    ServiceHi serviceHi;

    @RequestMapping("/feign/hi")
    public String hi() {
        return serviceHi.sayHiFromClient("service-feign");
    }
}
