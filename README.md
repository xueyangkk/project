## 关于

`go-gin-api` 是基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，约束项目组开发成员，规避混乱无序及自由随意的编码。

供参考学习，线上使用请谨慎！

集成组件：

1. 支持 [rate](https://golang.org/x/time/rate) 接口限流 
1. 支持 panic 异常时邮件通知 
1. 支持 [cors](https://github.com/rs/cors) 接口跨域 
1. 支持 [Prometheus](https://github.com/prometheus/client_golang) 指标记录 
1. 支持 [Swagger](https://github.com/swaggo/gin-swagger) 接口文档生成 
1. 支持 [GraphQL](https://github.com/99designs/gqlgen) 查询语言 
1. 支持 trace 项目内部链路追踪 
1. 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析 
1. 支持 [jwt](https://github.com/dgrijalva/jwt-go) 接口鉴权 
1. 支持 errno 统一定义错误码 
1. 支持 [zap](https://go.uber.org/zap) 日志收集 
1. 支持 [viper](https://github.com/spf13/viper) 配置文件解析
1. 支持 [gorm](https://gorm.io/gorm) 数据库组件
1. 支持 [go-redis](https://github.com/go-redis/redis/v7) 组件
1. 支持 RESTful API 返回值规范
1. 支持 gormgen、handlergen 代码生成工具
1. 支持 web 界面，使用的 [Light Year Admin 模板](https://gitee.com/yinqi/Light-Year-Admin-Using-Iframe)


## 文档索引

go-gin-api 文档由以下几个主要部分组成：

- 准备工作
- 快速开始
- 目录接口
- 核心封装
- 组件指南
- 工具包

**地址：[https://www.yuque.com/xinliangnote/go-gin-api/ngc3x5](https://www.yuque.com/xinliangnote/go-gin-api/ngc3x5)**

## 其他

查看 Jaeger 链路追踪代码，请查看 [v1.0版](https://exams-api/releases/tag/v1.0)，文档点这里 [jaeger.md](https://exams-api/blob/master/docs/jaeger.md) 。

## Special Thanks

[@koketama](https://github.com/koketama)

## Learning together

![](https://github.com/xinliangnote/Go/blob/master/00-基础语法/images/qr.jpg)

