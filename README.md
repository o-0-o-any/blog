# Gin+GORM简易博客项目

这是一个基于Go语言的Gin框架+GORM实现的简易博客网页项目，主要用于学习和实践Gin框架的路由、中间件、模板渲染，以及GORM的数据库操作等核心知识点。

## 项目介绍

本项目是学习完Gin和GORM后的实战练习项目，实现了博客的最基础功能，旨在巩固Go Web开发的核心知识点，熟悉Gin框架的使用和GORM对数据库的操作流程。

## 使用内容

- **Gin**
- **GORM**
- **Markdown文档编写README**
- **YAML文件与Viper读取**
- **Bcrypt哈希加密**
- **zap库搭建日志系统**
- **JWT双Token认证**
- **Session与Cookie的简单使用**
- **Docker部署**


## 项目结构
├── config/ # 配置文件读取（数据库配置、服务器配置等）\
├── routers/ # 路由注册（定义接口路由、绑定控制器）\
├── controller/ # 控制器（路由函数实现，处理请求和返回响应）\
├── middleware/ # 中间件（登录验证、跨域、日志等）\
├── models/ # 模型结构体（定义数据库表对应的结构体）\
├── db/ # 数据库初始化连接（GORM 连接数据库、迁移表结构）\
├── repository/ # 数据访问层（封装数据库 CURD 操作）\
├── utils/ # 工具函数（密码加密、数据校验等）\
├── templates/ # HTML 模板文件（前端页面模板）\
├── static/ # 静态文件（CSS、JS、图片等）\
├── go.mod # Go模块依赖\
└── go.sum # 依赖版本锁定\
- **前端代码主打能极简、能用就行，任何前端代码由豆包、DeepSeek、ChatGPT等AI工具生成，HTML、CSS等模板文件与静态文件都分布在templates文件夹下**