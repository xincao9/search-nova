package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sirupsen/logrus"
	"search-nova/controller/user"
	"search-nova/internal/authentication"
	"search-nova/internal/config"
	"search-nova/internal/constant"
	"search-nova/internal/logger"
	_ "search-nova/internal/manager"
	"net/http"
)

func main() {
	gin.SetMode(config.C.GetString(constant.ServerMode))
	engine := gin.New()
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{Output: logger.L.WriterLevel(logrus.DebugLevel)}), gin.RecoveryWithWriter(logger.L.WriterLevel(logrus.ErrorLevel)))
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
	user.Route(engine)
	authorized := engine.Group("/", authentication.Authentication)
	user.AuthenticationRoute(authorized)
	routeStatic(engine)
	addr := fmt.Sprintf(":%d", config.C.GetInt(constant.ServerPort))
	logger.L.Infof("Listening and serving HTTP on : %s", addr)
	if err := engine.Run(addr); err != nil {
		logger.L.Fatalf("Fatal error search-nova: %v\n", err)
	}
}

func routeStatic(engine *gin.Engine) {
	engine.Static("/assets", config.C.GetString(constant.AssetsRootDir))
	engine.Static("/js", config.C.GetString(constant.AssetsJsDir))
	engine.Static("/css", config.C.GetString(constant.AssetsCssDir))
	engine.Static("/img", config.C.GetString(constant.AssetsImgDir))
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/assets")
	})
}
