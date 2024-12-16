package config

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/exec"
	"search-nova/internal/constant"
	"strings"
)

var (
	C *viper.Viper
)

func init() {
	d := flag.Bool("d", false, "run app as a daemon with -d=true")
	c := flag.String("conf", "app.yaml", "configure file")
	if flag.Parsed() == false {
		flag.Parse()
	}
	// 后台进程执行
	if *d {
		args := os.Args[1:]
		for i := 0; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		err := cmd.Start()
		if err != nil {
			log.Fatalf("Fatal error conf : %v\n", err)
		}
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
	// 提取文件名
	for _, t := range []string{"yaml", "yml"} {
		if strings.HasSuffix(*c, t) {
			i := strings.LastIndex(*c, t)
			*c = string([]byte(*c)[:i-1])
		}
	}
	C = viper.New()
	C.SetConfigName(*c)
	C.SetConfigType("yaml")
	C.AddConfigPath("./conf")
	// 设置默认配置项
	C.SetDefault(constant.LoggerDir, "/tmp/search-nova/log")
	C.SetDefault(constant.LoggerLevel, "debug")
	C.SetDefault(constant.ServerMode, "debug")
	C.SetDefault(constant.ServerPort, 8080)
	C.SetDefault(constant.ServerCorsAccessControlAllowOrigin, "http://localhost:8081")
	C.SetDefault(constant.ManagerServerPort, 8090)
	C.SetDefault(constant.DataSource, "root:asdf@tcp(localhost:3306)/search_nova?charset=utf8&parseTime=true&loc=Local")
	C.SetDefault(constant.PublicDir, "./front/dist")
	C.SetDefault(constant.ElasticsearchAddresses, []string{"https://127.0.0.1:9200"})
	C.SetDefault(constant.ElasticsearchUsername, "elastic")
	C.SetDefault(constant.ElasticsearchPassword, "")
	C.SetDefault(constant.ElasticsearchIndex, "search_nova")
	err := C.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error conf : %v\n", err)
	}
}

func Route(engine *gin.Engine) {
	// 提供后台查看运行时配置的管理接口
	engine.GET("/config", func(c *gin.Context) {
		c.JSON(http.StatusOK, C.AllSettings())
	})
}
