package com.dkfjkjg.learn.springbootmybatis.mybatisinterceptors;

import org.apache.ibatis.executor.statement.StatementHandler;
import org.apache.ibatis.mapping.BoundSql;
import org.apache.ibatis.plugin.Interceptor;
import org.apache.ibatis.plugin.Intercepts;
import org.apache.ibatis.plugin.Invocation;
import org.apache.ibatis.plugin.Signature;
import org.apache.ibatis.reflection.DefaultReflectorFactory;
import org.apache.ibatis.reflection.MetaObject;
import org.apache.ibatis.reflection.ReflectorFactory;
import org.apache.ibatis.reflection.SystemMetaObject;
import org.springframework.stereotype.Component;

import java.util.Properties;

@Component
@Intercepts({
        @Signature(type = StatementHandler.class, method = "prepare",
        args = {java.sql.Connection.class, Integer.class})
})
public class UserTableNameInterceptor implements Interceptor {

    private static final ReflectorFactory DEFAULT_REFLECTOR_FACTORY = new DefaultReflectorFactory();

    @Override
    public Object intercept(Invocation invocation) throws Throwable {
        MetaObject metaObject = getMetaObject(invocation);
        BoundSql boundSql = (BoundSql) metaObject.getValue("delegate.boundSql");
        metaObject.setValue("delegate.boundSql.sql",
                boundSql.getSql().replace("t_user_prefix", "t_user"));
        System.out.println(invocation.getMethod());
        return invocation.proceed();
    }

    private MetaObject getMetaObject(Invocation invocation) {
        StatementHandler statementHandler = (StatementHandler) invocation.getTarget();
        // MetaObject是mybatis里面提供的一个工具类，类似反射的效果
        MetaObject metaObject = MetaObject.forObject(statementHandler,
                SystemMetaObject.DEFAULT_OBJECT_FACTORY,
                SystemMetaObject.DEFAULT_OBJECT_WRAPPER_FACTORY, DEFAULT_REFLECTOR_FACTORY);

        return metaObject;
    }

    @Override
    public Object plugin(Object target) {
        return Interceptor.super.plugin(target);
    }

    @Override
    public void setProperties(Properties properties) {
        Interceptor.super.setProperties(properties);
    }
}
