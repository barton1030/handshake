# handshake
简介：
该项目是一个异步交互中间件，用于多个项目之间解耦。
两个项目之间进行异步交互时需要创建相应的通信主题，用于数据收集、数据提取、异步回调、预警、熔断保护等等

目录架构：
app.             // 应用层
conduit          // 通信管道模块，用于管理不同层，不同模块单元的通信
  -- conduit.go  // 通信管道控制组建 
conf
  -- config.toml  // 项目配置文件、db、redis等配置信息
  -- parse.go     // 配置内容解析程序
domain            // 领域模块
helper            // 辅助模块 
  -- http.go      // http请求公共解析工具
  -- request.go   // http请求对象，支持POST、GET请求
persistent        // 持久层，用于项目实例数据的持久化
interface         // 接口层
middleware        // 中间件
engine            // topic执行和控制单元
service           // 业务服务层，实际项目逻辑层
main              // 项目启动入口
router            // 项目路由

使用步骤：
1、建表：
DROP TABLE IF EXISTS `hand_shake_role`;
CREATE TABLE `hand_shake_role` (
`id` int NOT NULL AUTO_INCREMENT,
`status` int DEFAULT '1',
`name` varchar(255) NOT NULL DEFAULT '',
`permission_map` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
`creator` int NOT NULL DEFAULT '0',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;


DROP TABLE IF EXISTS `hand_shake_user`;
CREATE TABLE `hand_shake_user` (
`id` int NOT NULL AUTO_INCREMENT,
`status` smallint DEFAULT '1',
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
`phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
`pwd` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
`role_id` int NOT NULL DEFAULT '0',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) USING BTREE,
KEY `phone` (`phone`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;


DROP TABLE IF EXISTS `hand_shake_topic`;
CREATE TABLE `hand_shake_topic` (
`id` int NOT NULL AUTO_INCREMENT,
`name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
`status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0',
`max_retry_count` int NOT NULL DEFAULT '1',
`min_concurrency` int NOT NULL DEFAULT '1',
`max_concurrency` int NOT NULL DEFAULT '1',
`fuse_salt` int NOT NULL DEFAULT '10',
`alarm` varchar(750) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
`callback` varchar(750) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
`creator` int NOT NULL DEFAULT '0',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`) USING BTREE,
KEY `name` (`name`) USING BTREE,
KEY `creator` (`creator`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;


DROP TABLE IF EXISTS `hand_shake_log`;
CREATE TABLE `hand_shake_log` (
`id` int NOT NULL AUTO_INCREMENT,
`data` varchar(750) NOT NULL DEFAULT '',
`business_id` int NOT NULL DEFAULT '0',
`business_type` smallint NOT NULL DEFAULT '0',
`creator` int NOT NULL DEFAULT '0',
`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `businessIdAndType` (`business_id`,`business_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

2、配置conf中数据库和redis的基础信息

3、插入基础数据：
insert into hand_shake_role (`status`, `name`, permission_map, creator) VALUES(1, "超级管理员", "{\"all\":true}", 0);
insert into hand_shake_user (`status`, `name`, phone, pwd, role_id) VALUES (1, "admin", "1888888888", "", 1);

4、添加用户和角色
  1、通过admin账号，可以添加任意数量的角色和用户信息，其中角色名称不能重复，用户手机号不能重复
  2、用户创建完成后，将用户id提供给对方，其创建topic、修改、废弃、启动和关闭等动作都需要传该参数
  3、可以通过调整角色的权限来改变相关用户的权限，甚用。
  4、admin账号拥有，角色和用户操作的全部权限。
  5、角色权限设置方式，路由：true/false, true代表有该路由下的逻辑执行权限，false代表无。

5、回调和预警逻辑因不同个人和公司的实现工具不同，故只完成基础逻辑，具体的实现逻辑，个人根据具体场景填充
   domain/topic/alarm.go
   domain/topic/callback.go

6、编译启动、通过topic/start