package impl

import (
	"context"
	"go-vblog/apps/blog"
)

func (i *Impl) save(ctx context.Context, instance *blog.Blog) error {
	return i.DB().WithContext(ctx).Create(instance).Error
}
