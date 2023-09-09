package main

import (
	"go-vblog/cmd"
)

func main() {
	//if err := conf.LoadConfigFromEnv(); err != nil {
	//	panic(err)
	//}
	//protocol.NewHTTP().Start()
	cmd.Execute()
}
