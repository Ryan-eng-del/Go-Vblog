package impl

import (
	"database/sql"
	"go-vblog/apps/blog"
	"go-vblog/conf"
	"gorm.io/gorm"
)

type Impl struct {
	db      *gorm.DB
	db_real *sql.DB
}

func NewBlogServiceImpl() *Impl {
	return &Impl{}
}

func (i *Impl) Name() string {
	return blog.AppName
}

func (i *Impl) DB() *gorm.DB {
	return i.db.Table(i.Name())
}

func (i *Impl) Init() error {
	i.db = conf.C().MySQL.GetORMDB().Debug()
	i.db_real = conf.C().MySQL.GetDB()
	return nil
}
