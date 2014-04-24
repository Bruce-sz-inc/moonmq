package main

import (
	"flag"
	"github.com/siddontang/moonmq/client"
)

var addr = flag.String("addr", "127.0.0.1:11181", "moonmq listen address")
var queue = flag.String("queue", "test_queue", "queue want to bind")

func main() {
	cfg := client.NewDefaultConfig()
	cfg.BrokerAddr = *addr

	c, err := client.NewClientWithConfig(cfg)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, err = c.PublishFanout(*queue, []byte("hello world"))
	if err != nil {
		panic(err)
	}
}
