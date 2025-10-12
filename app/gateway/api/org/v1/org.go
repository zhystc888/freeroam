package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta   `path:"/org/list" tags:"组织管理" method:"get" summary:"列表"`
	Name     string `json:"name" dc:"组织名，模糊搜索"`
	ParentId *int64 `json:"parentId" v:"integer|min:0#上级id必须是整数|上级id必须大于0" dc:"父id"`
	Code     string `json:"code" description:"组织编码"`
	Type     int64  `json:"type" v:"integer|min:0#组织分类不正确1|组织分类不正确2" dc:"组织分类"`
	PageReq
}

type GetListRes struct {
	List  []*GetListItem
	Total int64
}

type GetListItem struct {
	Id       int64  `json:"id" dc:"主键"`              // 主键
	Name     string `json:"name" dc:"组织名称"`        // 组织名称
	Code     string `json:"code" dc:"组织编码"`        // 组织编码
	Type     int64  `json:"type" dc:"组织分类，枚举"`   // 组织分类，枚举
	Status   int32  `json:"status" dc:"组织状态，枚举"` // 组织状态，枚举
	CreateAt string `json:"createAt" dc:"创建时间"`
}

type GetReq struct {
	g.Meta `path:"/org" tags:"组织管理" method:"get" summary:"查看组织详情"`
	Id     int64 `json:"id" dc:"组织id" v:"required|integer|min:1#ID不能为空|ID必须是整数|ID不能小于1"`
}

type GetRes struct {
	Id          int64         `json:"id" dc:"主键"`              // 主键
	ParentId    int64         `json:"parentId" dc:"父id"`        // 父id
	Name        string        `json:"name" dc:"组织名称"`        // 组织名称
	Code        string        `json:"code" dc:"组织编码"`        // 组织编码
	Type        int64         `json:"type" dc:"组织分类，枚举"`   // 组织分类，枚举
	Status      int64         `json:"status" dc:"组织状态，枚举"` // 组织状态，枚举
	Supervisors []*Supervisor `json:"supervisors" dc:"部门主管列表"`
}

type Supervisor struct {
	UserId int64  `json:"id" orm:"user_id" dc:"主管用户id"`
	Name   string `json:"name" orm:"name" dc:"主管用户姓名"`
}
