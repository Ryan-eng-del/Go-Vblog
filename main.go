package main

import (
	"go-vblog/apps/blog/impl"
	"go-vblog/conf"
)

func main() {
	conf.LoadConfigFromEnv()
	if err := impl.NewBlogServiceImpl().Init(); err != nil {
		panic(err)
	}
}
