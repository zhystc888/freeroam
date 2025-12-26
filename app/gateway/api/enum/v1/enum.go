package v1

import "github.com/gogf/gf/v2/frame/g"

type GetEnumListReq struct {
	g.Meta `path:"/system/enum" tags:"枚举管理" method:"get" summary:"获取枚举列表"`
	Type   []string `p:"type" v:"required|min-length:1#枚举类型不能为空|至少需要一个枚举类型"`
}

type GetEnumListRes map[string][]GetEnumListResItem

type GetEnumListResItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Sort  int    `json:"sort"`
}
