package org.example.springbootmybatisplus;

import com.baomidou.mybatisplus.core.toolkit.Wrappers;
import com.baomidou.mybatisplus.extension.toolkit.Db;
import com.baomidou.mybatisplus.extension.toolkit.SimpleQuery;
import org.example.springbootmybatisplus.entities.User;
import org.example.springbootmybatisplus.mapper.UserMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.util.Assert;

import java.util.List;

@SpringBootTest
class SpringbootMybatisplusApplicationTests {

    @Test
    void contextLoads() {
    }

    @Autowired
    private UserMapper userMapper;

    @Test
    public void testSelect() {
        System.out.println(("----- selectAll method test ------"));
        List<User> userList = userMapper.selectList(null);
        Assert.isTrue(5 == userList.size(), "");
        userList.forEach(System.out::println);
    }

    @Test
    public void testSelectById() {
        System.out.println(("----- testSelectById method test ------"));
        User user = userMapper.selectById(1);
        System.out.println(user);
    }

    @Test
    public void testSimpleQuery() {
        User user = Db.getOne(Wrappers.query(User.class).eq("id", 2));
        System.out.println(user);
    }

    @Test
    public void testLamdaQuery() {
        List<User> userList = Db.list(Wrappers.lambdaQuery(User.class).eq(User::getId, "3"));
        Assert.isTrue(1 == userList.size(), "");
        userList.forEach(System.out::println);
    }

}
