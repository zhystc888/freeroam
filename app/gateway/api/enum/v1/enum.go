package v1

import "github.com/gogf/gf/v2/frame/g"

type GetEnumListReq struct {
	g.Meta `path:"/enum/list" tags:"枚举管理" method:"post" summary:"获取枚举列表"`
	List   []*GetEnumListReqItem `json:"list" v:"required|min-length:1#枚举列表不能为空|枚举列表至少需要1个元素"`
}

type GetEnumListReqItem struct {
	Code      string `json:"code" dc:"枚举编码" v:"required#枚举编码不能为空"`
	TableName string `json:"tableName" dc:"表名" v:"required#表名不能为空"`
	Module    string `json:"module" dc:"模块"`
}

type GetEnumListRes struct {
	List map[string][]*GetEnumListResItem `json:"list"`
}

type GetEnumListResItem struct {
	Value string `json:"value" dc:"值"`
	Name  string `json:"name" dc:"名"`
}
