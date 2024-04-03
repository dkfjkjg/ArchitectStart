package org.example.configjdbcclient;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class ConfigJdbcClientApplication {

    public static void main(String[] args) {
        SpringApplication.run(ConfigJdbcClientApplication.class, args);
    }

    @Value("${foo}")
    String foo;
    @RequestMapping(value = "/foo")
    public String hi(){
        return foo;
    }

    @Bean
    public CommandLineRunner commandLineRunner() {
        return args -> System.out.println(foo);
    }
}
