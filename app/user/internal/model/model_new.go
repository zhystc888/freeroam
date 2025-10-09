package model

type PageReq struct {
	Page  int32 `p:"page" v:"min:0#分页号码错误" dc:"分页号码" d:"1"`
	Limit int32 `p:"limit" v:"max:100#分页数量最大100条" dc:"分页数量，最大100" d:"10"`
}

func (p *PageReq) GetLimit() int {
	return int(p.Limit)
}

func (p *PageReq) GetPage() int {
	return int(p.Page)
}

func (p *PageReq) GetOffset() int {
	return int((p.Page - 1) * p.Limit)
}
