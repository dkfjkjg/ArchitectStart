package org.example.serviceribbon.demos;

import org.example.serviceribbon.services.HelloService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
public class HelloController {

    @Autowired
    HelloService helloService;

    @RequestMapping("/ribbon/hi")
    public String hi() {
        return helloService.hiService( "ribbon");
    }
}
