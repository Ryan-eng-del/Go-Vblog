package impl

import (
	"context"
	"go-vblog/apps/blog"
	"go-vblog/conf"
	"testing"
)

var blogService *Impl

func TestCreateBlog(t *testing.T) {
	req := blog.NewCreateBlogRequest()

	req.TitleName = "Vblog1"
	req.Summary = "文章概要"
	req.Content = "develop vblog system"

	ins, err := blogService.CreateBlog(context.Background(), req)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(ins)
}

func init() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	blogService = NewBlogServiceImpl()
	blogService.Init()
}
