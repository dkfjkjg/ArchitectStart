package com.dkfjkjg.learn.springbootmybatis.controller;

import com.dkfjkjg.learn.springbootmybatis.entity.TUser;
import com.dkfjkjg.learn.springbootmybatis.mapper.TUserMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/user")
public class UserController {

    @Autowired
    TUserMapper userMapper;

    @GetMapping("/listAll")
    public List<TUser> allUser() {
        return userMapper.selectAll();
    }
}
