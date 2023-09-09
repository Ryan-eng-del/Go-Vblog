package impl_test

import (
	"context"
	"testing"

	"gitee.com/go-course/go8/projects/vblog/api/apps"
	"gitee.com/go-course/go8/projects/vblog/api/apps/tag"
	"gitee.com/go-course/go8/projects/vblog/api/conf"
	"github.com/infraboard/mcube/logger/zap"

	_ "gitee.com/go-course/go8/projects/vblog/api/apps/all"
)

// svc 实现tag.Service接口
var svc tag.Service

func TestAddTag(t *testing.T) {
	req := tag.NewAddTagRequest()
	req.AddTag(&tag.CreateTagRequest{
		BlogId: 5,
		Key:    "Language",
		Value:  "Golang",
	})
	set, err := svc.AddTag(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(set)
}

func TestQueryTag(t *testing.T) {
	req := tag.NewQueryTagRequest()
	req.BlogId = 5
	set, err := svc.QueryTag(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestRemoveTag(t *testing.T) {
	req := tag.NewRemoveTagRequest()
	req.AddTagId(1)
	set, err := svc.RemoveTag(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func init() {
	// 加载配置
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	zap.DevelopmentSetup()

	// IOC容器服务实例对象初始化
	if err := apps.Init(); err != nil {
		panic(err)
	}

	svc = apps.GetService(tag.AppName).(tag.Service)
}
