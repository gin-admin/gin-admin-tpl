# gin-admin-tpl

> gin-admin basic develop template.

## 特性

- 遵循 `RESTful API` 设计规范 & 基于接口的编程规范
- 基于 `GIN` 框架，提供了丰富的中间件支持（JWTAuth、CORS、RequestLogger、RequestRateLimiter、TraceID、CasbinEnforce、Recover、GZIP）
- 基于 `Gorm 2.0` 的数据库访问层 - 全功能 ORM
- 基于 `WIRE` 的依赖注入 -- 依赖注入本身的作用是解决了各个模块间层级依赖繁琐的初始化过程
- 基于 `Logrus & Context` 实现了日志输出，通过结合 Context 实现了统一的 TraceID/UserID 等关键字段的输出(同时支持日志钩子写入到`Gorm`)
- 基于 `Swaggo` 自动生成 `Swagger` 文档 -- 独立于接口的 mock 实现
- 基于 `net/http/httptest` 标准包实现了 API 的单元测试
- 基于 `go mod` 的依赖管理(国内源可使用：<https://goproxy.cn/>)

## 依赖工具

```bash
go get -u github.com/cosmtrek/air
go get -u github.com/google/wire/cmd/wire
go get -u github.com/swaggo/swag/cmd/swag
```

- [air](https://github.com/cosmtrek/air) -- Live reload for Go apps
- [wire](https://github.com/google/wire) -- Compile-time Dependency Injection for Go
- [swag](https://github.com/swaggo/swag) -- Automatically generate RESTful API documentation with Swagger 2.0 for Go.

## 依赖框架

- [Gin](https://gin-gonic.com/) -- The fastest full-featured web framework for Go.
- [GORM](https://gorm.io/) -- The fantastic ORM library for Golang
- [Casbin](https://casbin.org/) -- An authorization library that supports access control models like ACL, RBAC, ABAC in Golang
- [Wire](https://github.com/google/wire) -- Compile-time Dependency Injection for Go

## 快速开始

```bash
# 使用AIR工具运行
$ air

# OR 基于Makefile运行
$ make start

# OR 使用go命令运行
$ go run cmd/gin-admin/main.go web -c ./configs/config.toml
```

> 启动成功之后，可在浏览器中输入地址进行访问：[http://127.0.0.1:10099/swagger/index.html](http://127.0.0.1:10099/swagger/index.html)

## 生成`swagger`文档

```bash
# 基于Makefile
make swagger

# OR 使用swag命令
swag init --parseDependency --generalInfo ./cmd/${APP}/main.go --output ./internal/app/swagger
```

## 重新生成依赖注入文件

```bash
# 基于Makefile
make wire

# OR 使用wire命令
wire gen ./internal/app
```

## [gin-admin-cli](https://github.com/gin-admin/gin-admin-cli) 工具使用

### 创建项目

```bash
gin-admin-cli new -d test-gin-admin -p test-gin-admin -m
```

### 快速生成业务模块

#### 创建模板 task.yaml

```bash
name: Task
comment: 任务管理
fields:
  - name: Code
    type: string
    comment: 任务编号
    required: true
    binding_options: ""
    gorm_options: "size:50;index;"
  - name: Name
    type: string
    comment: 任务名称
    required: true
    binding_options: ""
    gorm_options: "size:50;index;"
  - name: Memo
    type: string
    comment: 任务备注
    required: false
    binding_options: ""
    gorm_options: "size:1024;"
```

#### 执行生成命令并运行

```bash
gin-admin-cli g -d . -p gin-admin-tpl -f ./task.yaml

# 生成 swagger
make swagger

# 生成依赖项
make wire

# 运行服务
make start
```

## 目录结构

```
├── cmd
│   └── gin-admin-tpl
│       └── main.go       # 入口文件
├── configs
│   ├── config.toml       # 配置文件
├── docs                  # 文档
├── internal
│   └── app
│       ├── api           # API 处理层
│       ├── config        # 配置文件映射
│       ├── contextx      # 统一上下文处理
│       ├── dao           # 数据访问层
│       ├── ginx          # gin 扩展模块
│       ├── middleware    # gin 中间件模块
│       ├── module        # 通用业务处理模块
│       ├── router        # 路由层
│       ├── schema        # 统一入参、出参对象映射
│       ├── service       # 业务逻辑层
│       ├── swagger       # swagger 生成文件
│       ├── test          # 模块单元测试
├── pkg
│   ├── auth
│   │   └── jwtauth       # jwt 认证模块
│   ├── errors            # 错误处理模块
│   ├── gormx             # gorm 扩展模块
│   ├── logger            # 日志模块
│   │   ├── hook
│   └── util              # 工具包
│       ├── conv
│       ├── hash
│       ├── json
│       ├── snowflake
│       ├── structure
│       ├── trace
│       ├── uuid
│       └── yaml
└── scripts               # 统一处理脚本
```
