package main

// 引用具体模块
import (
	"api-server/config"
	"api-server/model"
	v "api-server/pkg/version"
	"api-server/router"
	"api-server/router/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zxmrlc/log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

// 参数化信息，可以通过 ./exec -h 进行参数化提示
var (
	// 配置文件默认路径，默认为空，通过参数c指定。
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")

	// 获得版本信息
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	// 载入参数相关内容
	pflag.Parse()

	// 初始化版本参数 -v 进行参数输出
	// 通过make命令，从Makefile进行版本信息提取
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// 初始化配置文件参数，默认 ./config/config.yaml
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 初始化mysql
	model.DB.MInit()
	//初始化redis
	model.RD.RInit()

	//根据实际情况，初始化数据表
	model.DB.Self.AutoMigrate(&model.AssetModel{})
	model.DB.Self.AutoMigrate(&model.UserModel{})

	//如果表里不包含指定数据，进行插入
	if err := model.DB.Self.Where("assetid = ? ", "1111").First(&model.AssetModel{}).Error; err != nil {
		asset := &model.AssetModel{AssetID: "1111", Assetname: "bbbb"}
		model.DB.Self.Create(&asset)
	}

	// 延时关闭数据库
	defer model.DB.Close()

	// 从配置文件读取运行模式 release/test/debug
	gin.SetMode(viper.GetString("runmode"))

	// 启动gin实例
	g := gin.New()

	// Routes.
	router.Load(
		// gin引擎.
		g,

		// 中间件.
		middleware.Logging(),
		middleware.RequestId(),
		middleware.CheckCors(),
	)

	// 在程序启动的时候，进行心跳检查，确保服务已经启动.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// 从配置文件，读取相关内容并启动监听
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// 使用get方式获得服务器状态是否正常，确保服务真的起来了。
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// 每1秒测试一次
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)

		// 验证/health的返回是否为200
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
	}
	return errors.New("Cannot connect to the router.")

}
