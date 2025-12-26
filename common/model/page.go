package model

// PageReq 分页请求结构体
// 用于 HTTP API 和内部服务统一的分页参数
type PageReq struct {
	Page  int32 `json:"page" p:"page" v:"min:0#分页号码错误" dc:"分页号码" d:"1"`
	Limit int32 `json:"limit" p:"limit" v:"max:100#分页数量最大100条" dc:"分页数量，最大100" d:"10"`
}

// GetLimit 获取分页数量（转换为 int）
func (p *PageReq) GetLimit() int {
	if p.Limit <= 0 {
		return 10 // 默认值
	}
	return int(p.Limit)
}

// GetPage 获取分页号码（转换为 int）
func (p *PageReq) GetPage() int {
	if p.Page <= 0 {
		return 1 // 默认值
	}
	return int(p.Page)
}

// GetOffset 计算分页偏移量
func (p *PageReq) GetOffset() int {
	page := p.GetPage()
	limit := p.GetLimit()
	return (page - 1) * limit
}

