

## 项目骨架

```sh
├── Makefile # Makefile项目脚手架
├── README.md # 项目说明文档
├── apps # 业务模块, 管理着所有的业务服务
│   ├── README.md
│   ├── all # 业务模块加载模块
│   │   └── impl.go
│   ├── blog    # blog 服务模块
│   │   ├── README.md
│   │   ├── api
│   │   │   ├── README.md
│   │   │   ├── blog.go
│   │   │   ├── http.go
│   │   │   └── http_test.go
│   │   ├── const.go
│   │   ├── impl
│   │   │   ├── blog.go
│   │   │   ├── blog_test.go
│   │   │   ├── dao.go
│   │   │   └── impl.go
│   │   ├── interface.go
│   │   └── model.go
│   ├── ioc.go # ioc模块
│   ├── tag # tag 服务模块
│   │   ├── README.md
│   │   ├── api
│   │   │   ├── http.go
│   │   │   └── tag.go
│   │   ├── impl
│   │   │   ├── impl.go
│   │   │   ├── tag.go
│   │   │   └── tag_test.go
│   │   ├── interface.go
│   │   └── model.go
│   └── user
│       ├── README.md
│       ├── api
│       │   ├── auth.go
│       │   └── http.go
│       ├── interface.go
│       └── model.go
├── cmd # 项目CLI
│   ├── README.md
│   ├── root.go
│   └── start.go
├── conf # 项目配置管理模块
│   ├── README.md
│   ├── config.go
│   ├── config_test.go
│   └── load.go
├── docs # 项目详细文档
│   └── api
│       └── vblog.postman_collection.json
├── etc # 项目配置文件
│   ├── config.env
│   ├── config.toml
│   └── unit_test.env
├── go.mod
├── go.sum
├── main.go # 项目入口
├── protocol # 项目对外暴露接口时相关协议
│   ├── README.md
│   ├── auth # 协议认证中间件
│   │   └── basci.go
│   └── http.go # 基于HTTP协议的暴露
└── version #  项目版本管理
    └── version.go
```

## 工程配置对象管理

- 基于文件(json,ymal,toml), 基于 toml 的格式来作为程序的配置
    - json: json.Marshal 标准库
    - ymal: 第三方库
    - toml: 第三方库: "github.com/BurntSushi/toml"
    - env: 基于环境变量, 容器部署时很有用, 如果解析环境变量, os.GetEnv, "github.com/caarlos0/env/v6", 通过定义 Struct Tag 直接帮你完成 环境变量映射

## 非功能性功能开发

- CLI:
- protocol: HTTP Server
- Makefile
