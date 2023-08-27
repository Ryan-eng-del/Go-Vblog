package main

import (
	"go-vblog/conf"
	"go-vblog/protocol"
)

func main() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	protocol.NewHTTP().Start()

}
