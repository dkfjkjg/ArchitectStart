package org.example.springbootmybatisplus.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Select;
import org.example.springbootmybatisplus.entities.User;

public interface UserMapper extends BaseMapper<User> {

    @Select("select * from user where id = #{id}")
    User getById(Integer id);

}
