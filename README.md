# handshake
简介：
该项目是一个异步交互中间件，用于多个项目之间解耦使用。
两个项目之间进行异步交互时需要创建响应的通信主题，用于数据收集、数据提取、异步回调、预警、熔断保护等等

目录架构：
app.             // 应用层
  -- user        // 用户模块
  -- topic       // 主题模块
conf
  -- config.toml  // 项目配置文件、db、redis等配置信息
helper           // 辅助模块 
persistent       // 持久层，用于项目实例数据的持久化
engine           // topic执行和控制单元
service          // 业务服务层，实际项目逻辑层
    -- engine    // 引擎模块
    -- user      // 用户模块
    -- role      // 角色模块，权限管理模块
    -- topic     // 主题模块 
    -- internal  // 接口模块，用于模块直接的解耦
 main.           // 项目启动入口
    -- main.go
 router          // 项目路由
    -- router.go 
