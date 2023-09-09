package tag

import "encoding/json"

func NewTagSet() *TagSet {
	return &TagSet{
		Items: []*Tag{},
	}
}

type TagSet struct {
	Items []*Tag `json:"items"`
}

func (s *TagSet) String() string {
	dj, _ := json.Marshal(s)
	return string(dj)
}

func (s *TagSet) Add(item *Tag) {
	s.Items = append(s.Items, item)
	return
}

// 用于存储文章标签, 整个key, value组成一个tag
type Tag struct {
	// 标签Id
	Id int `json:"id"`
	// 创建时间: 用于排序
	CreateAt int64 `json:"create_at"`
	// Tag的具体数据
	*CreateTagRequest
}

type CreateTagRequest struct {
	// 关联的博客, 同一标签，允许打在不同博客上的
	BlogId int `json:"blog_id" validate:"required"`
	// 标签名称
	Key string `json:"key" validate:"required"`
	// 标签的value
	Value string `json:"value" validate:"required"`
	// 标签的颜色
	Color string `json:"color"`
}
