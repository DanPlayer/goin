# GOIN

### 简介
这是基于Gin的一款web框架，采用mvc架构，方便快捷开发使用

### 使用
```
go get github.com/DanPlayer/goin
```

### 文档

#### 下载依赖前设置代理（也可go get获取依赖包）
1. windows
   > set GOPROXY=https://goproxy.io
2. centos/macos
   > export GOPROXY=https://goproxy.io

#### 目录
```
- config     配置文件
  
- crontab    脚本文件  

- db         数据库配置  

- docs       swagger文档  

- middleware 中间件（权限、session）  

- models     模型层（数据库操作）  

- pojo       数据传输（对外数据管理）  

- rdb        redis管理层  

- routers    路由层  
   -- application    -- 路由应用
   --- v1               -- 版本（v1）  

- scripts    shell脚本  

- service    服务层  
   -- impl           -- 逻辑层  

- tests      单元测试  

- utils      工具层  
```

#### config 配置
```
详细配置见：

app.dev.yaml 测试环境配置

app.prod.yaml 生产环境配置

例子：go run main.go dev
```

#### swagger 独立服务器:
1. 安装swagger `go get -u github.com/go-swagger/go-swagger/cmd/swagger`
2. 生成文档 `swagger generate spec -o ./swagger.json`
3. 启动文档服务器

#### swagger 集成服务器
1. 文档地址[gin_swagger](https://github.com/swaggo/gin-swagger) 
2. 安装swag `go get -u github.com/swaggo/swag/cmd/swag`
3. 根目录执行 `swag init`
4. 访问地址 `http://localhost:8000/docs/index.html`

#### middleware 中间件
```
auth MakeMemberToken 生产验证密钥

auth ParseMemberToken 验证密钥

在router中使用
v1.GET("/index", middleware.MemberAuth(config.Auth.Token), helloV1.Hello)
```

#### service 服务
```
service 作为我们的服务层

我们在编写一个服务的时候，先添加服务接口

service/hello.go

// Service 服务
type Service interface {
	// Hello hello
	Hello()
}

再我们编写一个逻辑

service/impl/hello.go

type HelloImpl struct{}

func (u *HelloImpl) Hello() {
	fmt.Println("hello world")
	return


注册服务

service.go

import (
	"goin/service/hello"
	helloImpl "goin/service/hello/impl"
)

var HelloService hello.Service = new(helloImpl.HelloImpl)

```

#### model 模型使用
参考[GORM2.0](https://gorm.io/zh_CN/docs/index.html)

#### router 路由
```
routers/router.go里注册路由

如application路由

//application接口
application.Init(e, config)

在application/router.go里

func Hello(c *gin.Context) {
	utils.OutJsonOk(c, "这里是应用层服务，你好!")
}

func Init(e *gin.Engine, config *conf.AppConf) {
	hello := e.Group("/hello")
	{
		hello.GET("/", Hello)
		v1 := hello.Group("/v1")

		v1.GET("/index", middleware.MemberAuth(config.Auth.Token), helloV1.Hello)
	}
}

API管理
hello/v1/index.go
```

#### 启动服务 
`go run main.go`

### 构建
`go build -mod=vendor`