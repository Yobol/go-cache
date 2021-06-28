package main

import (
	"flag"
	"fmt"

	"github.com/yobol/go-cache/test/benchmark/client"
)

func main() {
	host := flag.String("h", "localhost", "cache server host")
	port := flag.String("p", "10616", "cache server port")
	op := flag.String("c", "get", "command, could be get/set/del")
	key := flag.String("k", "", "key")
	value := flag.String("v", "", "value")
	flag.Parse()
	remoteAddr := fmt.Sprintf("%s:%s", *host, *port)
	c := client.New(client.ClientTypeTCP, remoteAddr)
	cmd := &client.Cmd{*op, *key, *value, nil}
	c.Run(cmd)
	if cmd.Error != nil {
		fmt.Println("error:", cmd.Error)
	} else {
		fmt.Println(cmd.Value)
	}
}
