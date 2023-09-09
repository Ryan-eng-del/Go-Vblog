package apps

import "github.com/gin-gonic/gin"

var (
	// 保存着所有的对象
	services     = map[string]Service{}
	httpServices = map[string]HttpService{}
)

// Service 代表他是一个Ioc的Service Object
type Service interface {
	Init() error
	Name() string
}

// 作为一个实现了Gin HTTP Handler服务, 提供一个路由注册功能
type HttpService interface {
	Service
	Registry(gin.IRouter)
}

func RegistryHttp(svc HttpService) {
	httpServices[svc.Name()] = svc
}

// 初始化所有已经注册过来的实例
func Init() error {
	// 初始化 service
	for i := range services {
		if err := services[i].Init(); err != nil {
			return err
		}
	}

	return nil
}

// Registry 对象注册
func Registry(svc Service) {
	services[svc.Name()] = svc
}

// GetService 获取对象
func GetService(name string) any {
	if v, ok := services[name]; ok {
		return v
	}
	panic("service: " + name + " not registied")
}

func InitHttpService(rootRouter gin.IRouter) error {
	// 初始化 http service
	for i := range httpServices {
		api := httpServices[i]
		if err := api.Init(); err != nil {
			return err
		}
		// blogAPI.Registry(v1.Group("/blog"))
		api.Registry(rootRouter.Group(api.Name()))
	}

	return nil
}
