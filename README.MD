# 第四次作业
添加文件 webook/internal/repository/cache/local/code.go

问题：
1. 执行任务遍历删除的时候，正在校验key是否删除的时刻（这个key是过期的），新进来一个相同key，这个时候是否会删除这个key--（没测试出来）
2. 执行Verify里面的判断逻辑，这个需要改成mutex吗（是否存在，正在校验，key被重置了）

# 第三次作业 修改已有的部署方案
![img.png](picture/k8s详情.png)
![img.png](picture/k8s部署登录.png)


# 第二次作业
```sql
-- 初始化语句
CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` longtext COLLATE utf8mb4_general_ci,
  `ctime` bigint DEFAULT NULL,
  `utime` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
);

-- 新增语句
alter table users
    add nike_name VARCHAR(32) null after password,
    add birthday VARCHAR(16) null after nike_name,
    add signature VARCHAR(512) null after birthday;
```
- 用户编辑 users/edit
![编辑请求成功.png](picture/编辑请求成功.png)
![编辑昵称校验失败.png](picture/编辑昵称校验失败.png)
![编辑生日校验.png](picture/编辑生日校验.png)
![编辑个人简介校验.png](picture/编辑个人简介校验.png)

- 用户详情 users/profile
![用户详情.png](picture/用户详情.png)



