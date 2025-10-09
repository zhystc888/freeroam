package v1

import "github.com/gogf/gf/v2/frame/g"

type GetReq struct {
	g.Meta `path:"/member" tags:"成员" method:"get" summary:"查看成员详情"`
	UserID int64 `json:"userId" dc:"用户ID" v:"required|integer|min:1#用户ID不能为空|用户ID必须是整数|用户ID不能小于1"`
}

type GetRes struct {
	g.Meta `resEg:"resource/example/adminGetMember.json"`
	MemberBase
	Username string `json:"username" dc:"用户名"`
}
