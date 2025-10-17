package v1

type PageReq struct {
	Page  int32 `json:"page" v:"min:0#分页号码错误" dc:"分页号码" d:"1"`
	Limit int32 `json:"limit" v:"max:100#分页数量最大100条" dc:"分页数量，最大100" d:"10"`
}
