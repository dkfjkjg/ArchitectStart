package org.example.serviceribbon.demos;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
public class HelloController {

    @Autowired
    RestTemplate restTemplate;

    @RequestMapping("/ribbon/hi")

    public String hi() {
        return restTemplate.getForEntity("http://service-hi/hi?name=" + "ribbon", String.class).getBody();
    }
}
