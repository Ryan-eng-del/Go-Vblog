package blog

import (
	"bytes"
	"fmt"
	"strings"
)

type Status int

const (
	// STATUS_DRAFT 草稿
	STATUS_DRAFT Status = 0
	// STATUS_PUBLISHED 已发布
	STATUS_PUBLISHED Status = 1
)

// Status 实现 fmt Stringer接口
func (s Status) String() string {
	if v, ok := STATUS_MAP[s]; ok {
		return v
	}
	return fmt.Sprintf("%d", s)
}

// MarshalJSON Json允许用户自定义序列化和反序列化的逻辑
// 控制序列化的过程 ---> draft  {"status": "draft"}
func (s Status) MarshalJSON() ([]byte, error) {
	// 枚举为啥要全部转化为大写 Draft DRAFT draft
	// 字符串拼接 "draft"
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToLower(s.String()))
	b.WriteString(`"`)
	return []byte(b.String()), nil
}

// {"status": 1}
// 定义类型的反序列化 {"status": "published"}
func (s *Status) UnmarshalJSON(b []byte) error {
	ins, err := ParseStatusFromString(string(b))
	if err != nil {
		return err
	}
	// 需要把当前的Status赋值
	*s = ins
	return nil
}

// ParseStatusFromString ParseTagTypeFromString Parse TagType from string: "draft"
func ParseStatusFromString(str string) (Status, error) {
	key := strings.Trim(string(str), `"`)
	for k, v := range STATUS_MAP {
		if v == key {
			return k, nil
		}
	}

	return Status(-1), fmt.Errorf("unknown Status: %s", str)
}

var (
	STATUS_MAP = map[Status]string{
		STATUS_DRAFT:     "draft",
		STATUS_PUBLISHED: "published",
	}
)

// 更新模式
type UpdateMode string

const (
	// UPDATE_MODE_PUT 全量更新
	UPDATE_MODE_PUT UpdateMode = "put"
	// UPDATE_MODE_PATCH 部分更新
	UPDATE_MODE_PATCH UpdateMode = "patch"
)
