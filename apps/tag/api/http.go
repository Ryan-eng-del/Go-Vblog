package api

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	apps "go-vblog/apps/ioc"
	"go-vblog/apps/tag"
)

func NewHTTPAPI() *HTTPAPI {
	return &HTTPAPI{}
}

// HTTPAPI 定义用来 对外暴露HTTP服务，注册给协议Server(HTTP Server, Gin)
type HTTPAPI struct {
	service tag.Service
	log     logger.Logger
}

func (h *HTTPAPI) Init() error {
	h.service = apps.GetService(tag.AppName).(tag.Service)
	h.log = zap.L().Named("api.tag")
	return nil
}

func (h *HTTPAPI) Name() string {
	return tag.AppName
}

// URI 注册给 Gin
func (h *HTTPAPI) Registry(r gin.IRouter) {
	//r.Use(auth.BasicAuth)
	// 管理员接口, 都需要认证
	r.POST("/", h.AddTag)
	r.DELETE("/:id", h.RemoveTag)
}

func init() {
	apps.RegistryHttp(&HTTPAPI{})
}
