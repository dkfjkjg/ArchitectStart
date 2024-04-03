package com.dkfjkjg.learn.springbootredis.controller;

import com.dkfjkjg.learn.springbootredis.entity.TUser;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;

@RestController
@RequestMapping("/user")
public class UserController {

    @Resource
    private RedisTemplate<String, TUser> redisTemplate;

    @GetMapping("/add") // 浏览器测试
    public ResponseEntity<TUser> add(TUser user) {
        redisTemplate.opsForValue().set(user.getId(), user);
        return new ResponseEntity<TUser>(user, HttpStatus.OK);
    }

    @GetMapping("/get/{id}")
    public ResponseEntity<TUser> add(@PathVariable("id") String id) {
        TUser tUser = redisTemplate.opsForValue().get(id);
        return new ResponseEntity<TUser>(tUser, HttpStatus.OK);
    }
}
