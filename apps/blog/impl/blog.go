package impl

import (
	"context"
	"github.com/infraboard/mcube/exception"
	"go-vblog/apps/blog"
)

func (i *Impl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate create blog request error")
	}
	instance := blog.NewCreateBlog(req)

	if err := i.save(ctx, instance); err != nil {
		return nil, err
	}
	return instance, nil
}

// UpdateBlog 更新文章
func (*Impl) UpdateBlog(context.Context, *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

// DeleteBlog 文章的删除
// 为什么删除后，还要返回数据, 方便前端和事件总线使用
func (*Impl) DeleteBlog(context.Context, *blog.DeleteBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

// QueryBlog 文章列表
func (*Impl) QueryBlog(context.Context, *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	return nil, nil
}

// DescribeBlog 文章详情
func (*Impl) DescribeBlog(context.Context, *blog.DescribeBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

// UpdateBlogStatus 更新文章的状态
func (*Impl) UpdateBlogStatus(context.Context, *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}
