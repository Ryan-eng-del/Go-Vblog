package tag

import (
	"context"
)

type Service interface {
	// QueryTag 查询标签
	QueryTag(context.Context, *QueryTagRequest) (*Set, error)
	// AddTag 文章添加Tag
	AddTag(context.Context, *AddTagRequest) (*Set, error)
	// RemoveTag 文章移除Tag
	RemoveTag(context.Context, *RemoveTagRequest) (*Set, error)
}
