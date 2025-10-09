package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AdminGetMemberListDto struct {
	UserID int64  `p:"userId" dc:"用户ID" v:"integer#userId必须是整数"`
	Name   string `p:"name" dc:"【名称|用户名】模糊搜索"`
	Status *int32 `p:"status" dc:"状态，见枚举" v:"in:0,1,2#成员状态参数不正确"`
	PageReq
}

type AdminGetMemberListVo struct {
	g.Meta `orm:"table:bbk_admin_user"`
	AdminBase
	CreateAt *gtime.Time
	User     *UserBase `json:"user" orm:"with:id=user_id"`
}

type AdminGetMemberVo struct {
	g.Meta `orm:"table:bbk_admin_user"`
	AdminBase
	Username string `json:"username" orm:"username"`
}
