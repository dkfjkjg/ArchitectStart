package org.example.servicefeign.service;

import org.springframework.stereotype.Component;

@Component
public class ServiceHiHystrix implements ServiceHi {

    @Override
    public String sayHiFromClient(String name) {
        return "sorry " + name;
    }
}
