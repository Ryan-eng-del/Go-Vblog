package impl

import (
	"context"
	"go-vblog/apps/blog"
)

func (*Impl) CreateBlog(context.Context, *blog.CreateBlogRequest) (*blog.Blog, error) {
	return nil, nil
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
