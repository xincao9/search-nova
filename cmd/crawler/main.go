package main

import (
	"search-nova/internal/crawler"
	"search-nova/internal/shutdown"
)

func main() {
	crawler.C.Loop()
	shutdown.S.Await()
}
