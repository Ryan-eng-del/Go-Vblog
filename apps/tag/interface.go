package tag

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

const (
	AppName = "tag"
)

var (
	// 使用参考: https://raw.githubusercontent.com/go-playground/validator/master/_examples/simple/main.go
	validate = validator.New()
)

type Service interface {
	// 查询标签
	QueryTag(context.Context, *QueryTagRequest) (*TagSet, error)
	// 文章添加Tag
	AddTag(context.Context, *AddTagRequest) (*TagSet, error)
	// 文章移除Tag
	RemoveTag(context.Context, *RemoveTagRequest) (*TagSet, error)
}

func NewAddTagRequest() *AddTagRequest {
	return &AddTagRequest{
		Tags: []*CreateTagRequest{},
	}
}

func NewTagSetFromAddTagRequest(req *AddTagRequest) *TagSet {
	set := NewTagSet()
	for i := range req.Tags {
		set.Add(NewTagFromAddTagRequest(req.Tags[i]))
	}

	return set
}

func NewTagFromAddTagRequest(req *CreateTagRequest) *Tag {
	return &Tag{
		CreateAt:         time.Now().Unix(),
		CreateTagRequest: req,
	}
}

type AddTagRequest struct {
	// 一次可以添加多个Tag
	Tags []*CreateTagRequest
}

func (req *AddTagRequest) BlogIds() (ids []int) {
	bid := map[int]struct{}{}

	// 完成了blog id 的去重
	for i := range req.Tags {
		bid[req.Tags[i].BlogId] = struct{}{}
	}

	for k := range bid {
		ids = append(ids, k)
	}

	return
}

func (req *AddTagRequest) AddTag(tag *CreateTagRequest) {
	req.Tags = append(req.Tags, tag)
}

func (req *AddTagRequest) Validate() error {
	// for i := range req.Tags {
	// 	validate.Struct(req.Tags[i])
	// }
	return validate.Struct(req)
}

func NewRemoveTagRequest() *RemoveTagRequest {
	return &RemoveTagRequest{
		TagIds: []int{},
	}
}

type RemoveTagRequest struct {
	TagIds []int `json:"tag_ids"`
}

func (req *RemoveTagRequest) AddTagId(id int) {
	req.TagIds = append(req.TagIds, id)
}

func NewQueryTagRequest() *QueryTagRequest {
	return &QueryTagRequest{
		TagIds: []int{},
	}
}

type QueryTagRequest struct {
	TagIds []int `json:"tag_ids"`
	BlogId int   `json:"blog_id"`
}

func (req *QueryTagRequest) AddTagId(ids ...int) {
	req.TagIds = append(req.TagIds, ids...)
}
