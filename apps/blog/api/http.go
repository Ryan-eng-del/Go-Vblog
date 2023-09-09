package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go-vblog/apps/blog"
	apps "go-vblog/apps/ioc"
	"go-vblog/apps/tag"
)

type HTTPAPI struct {
	service blog.Service
	logger  logger.Logger
	tag     tag.Service
}

func NewHTTPAPI() *HTTPAPI {
	return &HTTPAPI{}
}

func (*HTTPAPI) Name() string {
	return blog.AppName
}

func (h *HTTPAPI) Init() error {
	h.service = apps.GetService(blog.AppName).(blog.Service)
	h.tag = apps.GetService(tag.AppName).(tag.Service)
	h.logger = zap.L().Named("api.blog")
	return nil
}

func (h *HTTPAPI) Registry(route gin.IRouter) {
	route.POST("/", h.CreateBlog)
	route.GET("/:id", h.DescribeBlog)
	route.DELETE("/:id", h.DeleteBlog)
	route.PUT("/:id", h.PutBlog)
	route.PATCH("/:id", h.PatchBlog)
}

func init() {
	apps.RegistryHttp(&HTTPAPI{})
}
