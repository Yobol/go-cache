package main

import (
	"github.com/yobol/go-cache/api"
	cache "github.com/yobol/go-cache/core"
)

func main() {
	c, err := cache.New(cache.CacheTypeMemory)
	if err != nil {
		panic(err)
	}
	api.New(c).Listen()
}
