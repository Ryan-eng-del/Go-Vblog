package impl

import (
	"context"
	"dario.cat/mergo"
	"fmt"
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

// QueryBlog 文章列表
func (i *Impl) QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	set := blog.NewBlogSet()
	query := i.DB()

	if req.Keywords != "" {
		query = query.Where("title_name LIKE ? OR content LIKE ?", "%"+req.Keywords+"%",
			"%"+req.Keywords+"%")
	}

	if req.Status != nil {
		query = query.Where("status = ?", *req.Status)
	}

	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}

	query = query.Offset(req.Offset()).Limit(req.PageSize).Order("create_at DESC")

	if err := query.WithContext(ctx).Scan(&set.Items).Error; err != nil {
		return nil, err
	}

	return set, nil
}

// DescribeBlog 文章详情
func (i *Impl) DescribeBlog(ctx context.Context, req *blog.DescribeBlogRequest) (*blog.Blog, error) {
	instance := blog.NewCreateBlog(blog.NewCreateBlogRequest())
	query := i.DB().Where("id = ?", req.Id)

	if err := query.Find(instance).Error; err != nil {
		return nil, err
	}

	// 注意处理404
	if instance.Id == 0 {
		return nil, exception.NewNotFound("blog %d not found", req.Id)
	}

	return instance, nil

}

// DeleteBlog 文章的删除
// 为什么删除后，还要返回数据, 方便前端和事件总线使用
func (i *Impl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) (*blog.Blog, error) {
	instance, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.Id))
	if err != nil {
		return nil, err
	}

	if err := i.DB().Delete(instance).Error; err != nil {
		return nil, err
	}
	return instance, err
}

// UpdateBlog 更新文章
func (i *Impl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	instance, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(req.BlogId))

	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		instance.CreateBlogRequest = req.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		if err := mergo.MergeWithOverwrite(instance.CreateBlogRequest, req.CreateBlogRequest); err != nil {
			return nil, err
		}
	default:
		return nil, exception.NewBadRequest("update mode not support %s", req.UpdateMode)
	}

	fmt.Println(instance.CreateBlogRequest)

	if err := instance.CreateBlogRequest.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update blog request error")
	}

	if err != i.DB().WithContext(ctx).Save(instance).Error {
		return nil, err
	}
	return instance, nil
}

// UpdateBlogStatus 更新文章的状态
func (*Impl) UpdateBlogStatus(context.Context, *blog.UpdateBlogStatusRequest) (*blog.Blog, error) {
	return nil, nil
}
