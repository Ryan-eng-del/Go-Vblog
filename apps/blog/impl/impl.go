package impl

import (
	"database/sql"
	"go-vblog/apps/blog"
	apps "go-vblog/apps/ioc"
	"go-vblog/apps/tag"
	"go-vblog/conf"
	"gorm.io/gorm"
)

type Impl struct {
	db      *gorm.DB
	db_real *sql.DB
	tag     tag.Service
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
	i.tag = apps.GetService(tag.AppName).(tag.Service)
	return nil
}

// import _ "gitee.com/go-course/go8/projects/vblog/api/apps/blog/impl"
func init() {
	apps.Registry(&Impl{})
}
