package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go-vblog/apps/blog"
	"go-vblog/apps/ioc"
	"go-vblog/conf"
	"gorm.io/gorm"
)

func NewImpl(blog blog.Service) *Impl {
	return &Impl{
		blog: blog,
	}
}

// 依赖MySQL连接, 能与MySQL交互
// 负责实现Blog Service
type Impl struct {
	db *gorm.DB

	// 依赖Blog service 需要校验Blog是否存在
	blog blog.Service
	log  logger.Logger
}

func (i *Impl) Name() string {
	return "tag"
}

func (i *Impl) DB() *gorm.DB {
	return i.db.Table(i.Name())
}

// 当这个对象初始化的，会获取该对象需要的依赖
// 需要db这个依赖, 从配置文件中获取
func (i *Impl) Init() error {
	i.db = conf.C().MySQL.GetORMDB().Debug()
	i.log = zap.L().Named("tag")

	// 动态的从Ioc层获取对象依赖
	i.blog = apps.GetService(blog.AppName).(blog.Service)
	return nil
}

func init() {
	apps.Registry(&Impl{})
}
