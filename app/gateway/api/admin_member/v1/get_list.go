package v1

import (
	"bbk/app/gateway/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type GetMemberListReq struct {
	g.Meta `path:"/member/list" tags:"成员" method:"get" sm:"查看成员列表"`
	model.AdminGetMemberListDto
}

type GetMemberListRes struct {
	g.Meta `resEg:"resource/example/adminGetMemberList.json"`
	List   []GetMemberListItem `json:"list" dc:"数据列表" minItems:"10" maxItems:"10"`
	Total  int64               `json:"total" dc:"数据总数"`
}

type GetMemberListItem struct {
	MemberBase
	LastTime string `json:"lastTime" dc:"最后登陆时间"`
	LastIp   string `json:"lastIp" dc:"最后登陆IP"`
	CreateAt string `json:"createAt" dc:"创建时间"`
}
