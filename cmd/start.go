package cmd

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
	"go-vblog/conf"
	"go-vblog/protocol"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	confType string
	confFile string
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 API Server",
	Long:  "启动 API Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		loadGlobal()
		// 加载配置
		if err := loadConfig(); err != nil {
			return err
		}
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		wg := sync.WaitGroup{}
		
		wg.Add(1)
		http := protocol.NewHTTP()

		go func() {
			defer wg.Done()
			// 启动一个Goroutine再后台, 处理来自操作系统的信号
			for v := range ch {
				zap.L().Infof("receive signal: %s, stop service", v)
				switch v {
				case syscall.SIGHUP:
					if err := loadConfig(); err != nil {
						zap.L().Errorf("reload config error, %s", err)
					}
				default:
					// 优雅关闭HTTP 服务
					ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
					defer cancel()
					http.Stop(ctx)
				}
				return
			}
		}()

		if err := http.Start(); err != nil {
			return err
		}
		wg.Wait()
		return nil
	},
}

func loadGlobal() {
	// 全局日志对象
	zap.DevelopmentSetup()
}

func loadConfig() error {
	switch confType {
	case "env":
		return conf.LoadConfigFromEnv()
	case "file":
		return conf.LoadConfigFromToml(confFile)
	default:
		return fmt.Errorf("not supported config type, %s", confType)
	}
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	StartCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/config.toml", "the service config from file")
	RootCommand.AddCommand(StartCmd)
}
