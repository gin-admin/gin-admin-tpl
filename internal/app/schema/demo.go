package schema

import "time"

// Demo DemoManage对象
type Demo struct {
	ID        uint64    `json:"id"`                      // 唯一标识
	Code      string    `json:"code" binding:"required"` // 编号
	Name      string    `json:"name" binding:"required"` // 名称
	Memo      string    `json:"memo"`                    // 备注
	Status    int       `json:"status"`                  // 状态(1:启用 2:禁用)
	Creator   uint64    `json:"creator"`                 // 创建者
	CreatedAt time.Time `json:"created_at"`              // 创建时间
	UpdatedAt time.Time `json:"updated_at"`              // 更新时间
}

// DemoQueryParam 查询条件
type DemoQueryParam struct {
	PaginationParam
}

// DemoQueryOptions 查询可选参数项
type DemoQueryOptions struct {
	OrderFields  []*OrderField // 排序字段
	SelectFields []string      // 查询字段
}

// DemoQueryResult 查询结果
type DemoQueryResult struct {
	Data       Demos
	PageResult *PaginationResult
}

// Demos DemoManage列表
type Demos []*Demo
