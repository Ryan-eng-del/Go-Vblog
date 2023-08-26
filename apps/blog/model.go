package blog

import (
	"encoding/json"
	"go-vblog/apps/tag"
	"time"
)
import "github.com/go-playground/validator/v10"

var validate = validator.New()

func (b *Blog) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}

func NewCreateBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		CreateAt:          time.Now().Unix(),
		CreateBlogRequest: req,
		Status:            STATUS_DRAFT,
		Tags:              []*tag.Tag{},
	}
}

type Blog struct {
	// 文章Id
	Id int `json:"id" gorm:"primaryKey"`
	// 创建时间
	CreateAt int64 `json:"create_at"`
	// 更新时间
	UpdateAt int64 `json:"update_at"`
	// 发布时间
	PublishAt int64 `json:"publish_at"`
	// 用户提交数据
	*CreateBlogRequest
	// 文章状态 草稿/发布
	Status Status `json:"status"`
	// 博客标签
	Tags []*tag.Tag `json:"tags"`
}

type BlogSet struct {
	blogs []*Blog
	total int32
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{}
}

func (c *CreateBlogRequest) Validate() error {
	return validate.Struct(c)
}

type CreateBlogRequest struct {
	// 文章摘要信息,通过提前Content内容获取
	Summary string `json:"summary" validate:"required" gorm:"-"`
	// 文章图片
	TitleImg string `json:"title_img"`
	// 文章标题
	TitleName string `json:"title_name" validate:"required"`
	// 文章副标题
	SubTitle string `json:"sub_title"`
	// 文章内容
	Content string `json:"content" validate:"required"`
	// 文章作者
	Author string `json:"author"`
}

type UpdateBlogRequest struct {
	BlogId     int
	UpdateMode UpdateMode
	*CreateBlogRequest
}

type DeleteBlogRequest struct {
	Id int
}

type QueryBlogRequest struct {
	PageSize   int
	PageNumber int
	Keywords   string
	// 补充状态过滤参数, 用于web 前台 过滤已经发布的文章
	// 比如过滤 状态为发布的文章
	Status  *Status
	BlogIds []int
}

type DescribeBlogRequest struct {
	Id int
}

type UpdateBlogStatusRequest struct {
	// 文章Id
	Id int
	// 文章状态 草稿/发布
	Status Status
}
