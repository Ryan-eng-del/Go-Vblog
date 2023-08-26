package blog

type Status int

const (
	// STATUS_DRAFT 草稿
	STATUS_DRAFT Status = 0
	// STATUS_PUBLISHED 已发布
	STATUS_PUBLISHED Status = 1
)

// 更新模式

type UpdateMode string

const (
	// UPDATE_MODE_PUT 全量更新
	UPDATE_MODE_PUT UpdateMode = "put"
	// UPDATE_MODE_PATCH 部分更新
	UPDATE_MODE_PATCH UpdateMode = "patch"
)
