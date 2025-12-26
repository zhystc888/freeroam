package model

// ListRes 列表响应结构体（泛型）
// 用于统一所有列表查询的响应格式
type ListRes[T any] struct {
	List  []*T  `json:"list" dc:"数据列表"`
	Total int64 `json:"total" dc:"数据总数"`
}

// NewListRes 创建列表响应
func NewListRes[T any](list []*T, total int64) *ListRes[T] {
	return &ListRes[T]{
		List:  list,
		Total: total,
	}
}

// EmptyListRes 创建空列表响应
func EmptyListRes[T any]() *ListRes[T] {
	return &ListRes[T]{
		List:  make([]*T, 0),
		Total: 0,
	}
}

