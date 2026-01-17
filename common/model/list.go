package model

// ListRes 列表响应结构体（泛型）
// 用于统一所有列表查询的响应格式
type ListRes[T any] struct {
	List     []*T  `json:"list" dc:"数据列表"`
	Total    int64 `json:"total" dc:"数据总数"`
	Page     int64 `json:"page" dc:"页码"`
	PageSize int64 `json:"pageSize" dc:"每页数量"`
}

// NewListRes 创建列表响应
func NewListRes[T any](list []*T, total, page, pageSize int64) *ListRes[T] {
	return &ListRes[T]{
		List:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}
}

// EmptyListRes 创建空列表响应
func EmptyListRes[T any](page, pageSize int64) *ListRes[T] {
	return &ListRes[T]{
		List:     make([]*T, 0),
		Total:    0,
		Page:     page,
		PageSize: pageSize,
	}
}
