package impl

import (
	"database/sql"
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

func (i *Impl) Init() error {
	i.db = conf.C().MySQL.GetORMDB().Debug()
	i.db_real = conf.C().MySQL.GetDB()
	return nil
}
