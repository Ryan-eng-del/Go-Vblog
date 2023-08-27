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

func TestQueryBlog(t *testing.T) {
	queryRequest := &blog.QueryBlogRequest{}
	queryRequest.PageNumber = 1
	queryRequest.PageSize = 10
	//queryRequest.Keywords = "88"
	set, err := blogService.QueryBlog(context.Background(), queryRequest)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeBlog(t *testing.T) {
	request := blog.NewDescribeBlogRequest(1)
	instance, err := blogService.DescribeBlog(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(instance)
}

func TestDeleteBlog(t *testing.T) {
	request := blog.NewDeleteBlogRequest(1)
	instance, err := blogService.DeleteBlog(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(instance)
}
func TestUpdateBlog(t *testing.T) {
	request := blog.NewPatchUpdateBlogRequest(2)
	request.TitleName = "Vblog2"
	instance, err := blogService.UpdateBlog(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(instance)
}

func init() {
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	blogService = NewBlogServiceImpl()
	blogService.Init()
}
