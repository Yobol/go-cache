package main

import (
	"github.com/yobol/go-cache/api/http"
	"github.com/yobol/go-cache/api/tcp"
	cache "github.com/yobol/go-cache/core"
)

func main() {
	c, err := cache.New(cache.CacheTypeMemory)
	if err != nil {
		panic(err)
	}
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
