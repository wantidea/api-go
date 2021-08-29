## 使用 GO 搭建的 API 后台框架

一个适合新手上手的 API 框架，基于 Gin 开发。

环境需要 **RabbitMQ**，**Redis**，**Mysql**，开启 **go-module**

版本要求，无列出即无特殊要求

**Go >= go1.15.5**

**MySQL >= 8.0**



## 目录结构

```
.
|-- app 应用目录
|   |-- controllers 	控制器目录
|   |   |-- business			B端控制器
|   |   |   `-- top					总站，多个后台区分
|   |   |       `-- user 			用户模块
|   |   |           `-- user.go 	用户控制器
|   |   |
|   |   `-- client				C端控制器
|   |       `-- blog 				博客端，多个客户端区分
|   |           |-- article.go		文章模块控制器
|   |           `-- setting.go		设置模块控制器
|   |
|   |-- middlewares		中间件目录
|   |   |-- auth.go					认证中间件
|   |   |-- cors.go					跨域中间件
|   |   `-- recover.go				全局错误捕获中间件
|   |
|   |-- models			模型目录（目录区分模块，细分表名）
|   |   `-- user					用户模块
|   |       `-- user.go				用户表
|   |
|   |-- requests		请求参数验证目录
|   |   |-- business			B端请求
|   |   |   `-- top					总站，多个后台区分
|   |   |       |-- blog			博客模块
|   |   |       |   |-- article.go		 	文章请求
|   |   |       |   `-- setting.go		 	设置请求
|   |   |       `-- user			用户模块
|   |   |           |-- auth.go			 	认证请求
|   |   |           `-- user.go			 	用户请求
|   |   |
|   |   `-- client				C端请求
|   |       `-- blog				博客端，多个客户端区分
|   |           `-- article.go		文章模块请求
|   |
|   `-- services		逻辑层目录（模块划分）
|       |-- blog					博客模块
|       |   `-- article					文章-模块逻辑
|       |       `-- article.go			文章-逻辑
|       `-- upload					上传模块
|           `-- image.go				上传图片逻辑
|
|-- config				配置目录
|   |-- app.ini						主应用设置
|   |-- ...
|   `-- redis.ini					Redis 设置
|
|-- lib					扩展目录（自定义扩展）
|   |-- mongodb				mongodb 扩展
|   |   |-- client.go
|   |   `-- model.go
|   `-- orm					orm 扩展
|       |-- client.go
|       `-- model.go
|   
|-- routers				路由目录
|   |-- business				B端路由
|   |   `-- top						总站，多个后台区分
|   |       `-- user.go
|   |-- client					C端路由
|   |   `-- blog.go					博客端，多个客户端区分
|   `-- router.go
`-- runtime				应用运行时目录
    |-- logs					日志
    |   |-- 20210820.log
    |   `-- 20210824.log
    `-- upload					上传目录
        `-- images					图片上传目录
            `-- fe402327c0cc60f5ae1a055a467b3ebb.png
```

从目录接口可以看出 B端与C端 目录的区别：

以控制器为例

B端 站点->模块->具体控制器

C端 站点->模块即控制器

个人偏好这么设计是因为我所接触的 C端 的接口往往没有 B端 那么多与复杂，具体见仁见智，可自由更改。



## 应用

**B端 business**

1、top 总站 （管理所有后台的总站，如需要多个后台新建立目录即可）前端代码 **[backend-vue](https://github.com/wantidea/backend-vue)**

**C端 client**

1、blog 博客 （博客 Web 端），前端代码 **[blog-vue](https://github.com/wantidea/blog-vue)**



## 账号

top 总站： 管理员账号（admin） 密码（123456789）



## 数据库

项目目录下的 api_go.sql 文件



## 使用

1、部署完成开头所述的环境。

2、倒入数据库 api_go.sql

3、运行程序

```bash
go run main.go
```

