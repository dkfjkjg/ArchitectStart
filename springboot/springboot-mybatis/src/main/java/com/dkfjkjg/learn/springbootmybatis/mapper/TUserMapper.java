package com.dkfjkjg.learn.springbootmybatis.mapper;

import com.dkfjkjg.learn.springbootmybatis.entity.TUser;

import java.util.List;

public interface TUserMapper {
    int deleteByPrimaryKey(Integer id);

    int insert(TUser record);

    int insertSelective(TUser record);

    TUser selectByPrimaryKey(Integer id);

    int updateByPrimaryKeySelective(TUser record);

    int updateByPrimaryKey(TUser record);

    List<TUser> selectAll();
}