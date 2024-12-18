package main

import (
	"search-nova/internal/crawler"
	"search-nova/internal/logger"
	"search-nova/internal/shutdown"
)

func main() {
	logger.L.Infof("crawler 服务开始启动")
	crawler.C.Run()
	shutdown.S.Await()
	logger.L.Infof("crawler 服务关闭")
}
