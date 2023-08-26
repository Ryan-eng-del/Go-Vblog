package main

import (
	"go-vblog/apps/blog/impl"
	"go-vblog/conf"
)

func main() {
	conf.LoadConfigFromEnv()
	impl.NewBlogServiceImpl().Init()
}
