package model

// PageReq 分页请求结构体
// 用于 HTTP API 和内部服务统一的分页参数
type PageReq struct {
	Page     int64 `json:"page" p:"page" v:"min:1#页码必须大于0" dc:"分页号码" d:"1"`
	PageSize int64 `json:"pageSize" p:"pageSize" v:"min:1|max:200#每页数量必须大于0|每页数量不能超过200" dc:"分页数量，最大200" d:"10"`
}

// GetPage 获取分页号码
func (p *PageReq) GetPage() int64 {
	if p.Page <= 0 {
		return 1 // 默认值
	}
	return p.Page
}

// GetPageSize 获取分页数量
func (p *PageReq) GetPageSize() int64 {
	if p.PageSize <= 0 {
		return 10 // 默认值
	}
	if p.PageSize > 200 {
		return 200
	}
	return p.PageSize
}
