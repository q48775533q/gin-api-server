# 生成文件
```bash
# 生成的文件在$GOPATH/src/api-server/bin文件夹

# 注意修改对应环境变量

cd $GOPATH
git clone xxxxx

cd $GOPATH/src/api-server

go env -w GOPROXY=https://goproxy.cn,direct
go env -w GO111MODULE=on

go mod init
go mod tidy

make

cd $GOPATH/src/api-server/bin
./api.server 
```


# 软件及版本/模块对应版本

```
Gin                                     // web框架
github.com/dgrijalva/jwt-go             // jwt鉴权
github.com/spf13/viper                  // 配置文件读取
golang.org/x/crypto/bcrypt              // 加密
gopkg.in/go-playground/validator.v9     // 字段验证
github.com/jinzhu/gorm                  // mysql
github.com/jinzhu/gorm/dialects/mysql   // mysql
github.com/zxmrlc/log                  // Log相关
github.com/zxmrlc/log/lager            // Log相关
github.com/fsnotify/fsnotify            // 配置文件读取监听
github.com/redis/go-redis/v9            // redis https://redis.uptrace.dev/zh/guide/
```

# 目录结构说明
```
├── Makefile            // 用作编译生成文件，通过git的tag和gitCommit进行生成版本信息
├── conf                // 配置文件目录，证书目录
│    └── config.yaml
├── config              // 解析配置文件的go package
│    └── config.go
├── controller          // controller文件夹
│    ├── asset          // 假装是资产管理Demo
│    │    ├── asset.go  // 存放用户controller公用的函数、结构体等
│    │    ├── create.go
│    │    └── list.go
│    ├── user           // 假装是用户管理Demo
│    │   ├── create.go
│    │   ├── delete.go
│    │   ├── get.go
│    │   ├── list.go
│    │   ├── login.go   //用户登陆放在这里了
│    │   ├── update.go
│    │   └── user.go    // 存放用户controller公用的函数、结构体等
│    ├── conftroller.go
│    └── sd             // 健康检查
│         └── check.go
├── go.mod
├── main.go             // 入口文件
├── model               // # 数据库相关的操作统一放在这里，包括数据库初始化和对表的增删改查
│    ├── asset.go
│    ├── init.go        // 初始化和连接数据库
│    ├── model.go       // 存放一些公用的go struct
│    └── user.go
├── pkg                 //引用包
│    ├── auth           // 鉴权
│    │    └── auth.go
│    ├── constvar       // 常量
│    │    └── constvar.go
│    ├── errno          // 错误码
│    │    ├── code.go
│    │    └── errno.go
│    ├── token          // token认证
│    │    └── token.go
│    └── version        // 版本信息，配合makefile文件
│        ├── base.go
│        └── version.go
├── readme.md
├── router              // 路由和中间件
│    ├── middleware
│    │    ├── auth.go
│    │    ├── header.go
│    │    ├── logging.go
│    │    └── requestid.go
│    └── router.go
├── service             // 实际业务处理函数存放位置
│    ├── asset.go
│    └── user.go
└── util                // 工具类函数存放目录
     ├── util.go
     ├── redis.go       // 用作读取jwt相关内容，去token.go参考，本想做黑白名单，但是有违jwt的理念，暂时不管他 
     └── util_test.go
```