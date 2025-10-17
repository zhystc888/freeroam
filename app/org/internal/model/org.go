package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type GetOrgDto struct {
	Id int64 `p:"id" description:"组织id"`
}

type OrgListDto struct {
	Name     string `description:"组织名，模糊搜索"`
	ParentId *int64 `p:"parent_id"`
	Code     string `description:"组织编码"`
	Type     int64  `description:"组织分类"`
	*PageReq
}

type OrgVo struct {
	g.Meta      `orm:"table:free_org_structure"`
	Id          int64            `json:"id"        orm:"id"         description:"主键"`      // 主键
	ParentId    int64            `json:"parentId"  orm:"parent_id"  description:"父id"`     // 父id
	Name        string           `json:"name"      orm:"name"       description:"组织名称"`    // 组织名称
	Code        string           `json:"code"      orm:"code"       description:"组织编码"`    // 组织编码
	Type        int64            `json:"type"      orm:"type"       description:"组织分类，枚举"` // 组织分类，枚举
	Status      int64            `json:"status"    orm:"status"     description:"组织状态，枚举"` // 组织状态，枚举
	Supervisors []*OrgSupervisor `json:"supervisors" orm:"with:org_id=id"`
}

type OrgListVo struct {
	g.Meta   `orm:"table:free_org_structure"`
	Id       int64     `json:"id"        orm:"id"         description:"主键"`      // 主键
	Name     string    `json:"name"      orm:"name"       description:"组织名称"`    // 组织名称
	Code     string    `json:"code"      orm:"code"       description:"组织编码"`    // 组织编码
	Type     int64     `json:"type"      orm:"type"       description:"组织分类，枚举"` // 组织分类，枚举
	Status   int32     `json:"status"    orm:"status"     description:"组织状态，枚举"` // 组织状态，枚举
	CreateAt time.Time `json:"createAt" orm:"create_at" description:"创建时间"`
}

type OrgSupervisor struct {
	g.Meta `orm:"table:free_org_supervisor"`
	UserId int64       `json:"id" orm:"user_id" description:"主管用户id"`
	User   *UserMember `json:"user" orm:"with:user_id=user_id"`
}

type UserMember struct {
	g.Meta `orm:"table:free_user_member"`
	UserId int64  `json:"id" orm:"user_id" description:"主管用户id"`
	Name   string `json:"name" orm:"name" description:"主管用户姓名"`
}
