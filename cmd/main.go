package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
	"search-nova/controller/user"
	"search-nova/internal/authentication"
	"search-nova/internal/config"
	"search-nova/internal/constant"
	"search-nova/internal/logger"
	_ "search-nova/internal/manager"
)

func main() {
	gin.SetMode(config.C.GetString(constant.ServerMode))
	engine := gin.New()
	// http 请求日志记录以及500错误记录
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: logger.L.WriterLevel(logrus.InfoLevel)}))
	engine.Use(gin.RecoveryWithWriter(logger.L.WriterLevel(logrus.ErrorLevel)))
	// cors
	engine.Use(cors.New(cors.Options{
		AllowedOrigins: []string{config.C.GetString(constant.ServerCorsAccessControlAllowOrigin)},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))
	// 不需要认证的路由
	user.Route(engine)
	routeStatic(engine)
	// 需要认证的路由
	authorized := engine.Group("/", authentication.Authentication)
	user.AuthenticationRoute(authorized)
	// 启动
	addr := fmt.Sprintf(":%d", config.C.GetInt(constant.ServerPort))
	logger.L.Infof("Listening and serving HTTP on : %s", addr)
	if err := engine.Run(addr); err != nil {
		logger.L.Fatalf("Fatal error search-nova: %v\n", err)
	}
}

func routeStatic(engine *gin.Engine) {
	ar := config.C.GetString(constant.AssetsRootDir)
	engine.Static("/assets", ar)
	engine.Static("/js", filepath.Join(ar, "js"))
	engine.Static("/css", filepath.Join(ar, "css"))
	engine.Static("/img", filepath.Join(ar, "img"))
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/assets")
	})
}
