package tag

// Tag 用于存储文章标签, 整个key, value组成一个tag
type Tag struct {
	// 标签Id
	Id int `json:"id"`
	// 创建时间: 用于排序
	CreateAt int64 `json:"create_at"`
	// Tag的具体数据
	*CreateTagRequest
}

type TagSet struct {
	Items []*Tag `json:"items"`
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

type AddTagRequest struct {
	// 一次可以添加多个Tag
	Tags []*CreateTagRequest
}

type RemoveTagRequest struct {
	TagIds []int `json:"tag_ids"`
}
type QueryTagRequest struct {
	TagIds []int `json:"tag_ids"`
	BlogId int   `json:"blog_id"`
}
