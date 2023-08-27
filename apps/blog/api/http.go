package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go-vblog/apps/blog"
	"go-vblog/apps/blog/impl"
)

type HTTPAPI struct {
	service blog.Service
	logger  logger.Logger
}

func NewHTTPAPI() *HTTPAPI {
	return &HTTPAPI{}
}

func (*HTTPAPI) Name() string {
	return blog.AppName
}

func (h *HTTPAPI) Init() error {
	h.service = impl.NewBlogServiceImpl()
	h.service.(*impl.Impl).Init()
	zap.DevelopmentSetup()
	h.logger = zap.L().Named("api blog")
	return nil
}

func (h *HTTPAPI) Registry(route gin.IRouter) error {
	route.POST("/", h.CreateBlog)
	route.GET("/:id", h.DescribeBlog)
	route.DELETE("/:id", h.DeleteBlog)
	route.PUT("/:id", h.PutBlog)
	route.PATCH("/:id", h.PatchBlog)
	return nil
}
