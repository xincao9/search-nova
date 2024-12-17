package main

import (
	"search-nova/internal/crawler"
	"search-nova/internal/shutdown"
)

func main() {
	crawler.C.Run()
	shutdown.S.Await()
}
