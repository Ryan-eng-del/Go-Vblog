package protocol

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go-vblog/apps/blog/api"
	"go-vblog/conf"
	"net/http"
	"time"
)

func NewHTTP() *HTTP {
	r := gin.Default()
	//zap.DevelopmentSetup()

	return &HTTP{
		log:    zap.L().Named("server.http"),
		router: r,
		server: &http.Server{
			ReadHeaderTimeout: 60 * time.Second,
			ReadTimeout:       60 * time.Second,
			WriteTimeout:      60 * time.Second,
			IdleTimeout:       60 * time.Second,
			MaxHeaderBytes:    1 << 20, // 1M
			// 处理 HTTP 协议的逻辑, HTTP 中间件, 是一个路由(框架,Gin, ...)与处理(自己)
			Handler: r,
			Addr:    conf.C().App.HTTP.Addr(),
		},
	}
}

// HTTP 服务对象
type HTTP struct {
	router *gin.Engine
	server *http.Server
	log    logger.Logger
}

func (h *HTTP) Start() error {
	httpApi := api.HTTPAPI{}
	httpApi.Init()
	httpApi.Registry(h.router.Group("/vblog/api/v1"))
	h.log.Infof("http server serve on: %s", h.server.Addr)

	if err := h.server.ListenAndServe(); err != nil {
		// 处理正常退出情况
		if err == http.ErrServerClosed {
			return nil
		}
		return fmt.Errorf("server ListenAndServe error, %s", err)
	}

	return nil
}

func (h *HTTP) Stop(ctx context.Context) {
	h.log.Infof("server garceful shutdown ...")
	// HTTP Server 优雅关闭
	// 支持ctx, 10分 请求都没退出, 做超时设置
	if err := h.server.Shutdown(ctx); err != nil {
		h.log.Warnf("shutdown error, %s", err)
	} else {
		h.log.Infof("server garceful shutdown ok")
	}
}
