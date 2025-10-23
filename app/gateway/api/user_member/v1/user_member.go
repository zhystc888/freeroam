package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta   `path:"/user_member/list" tags:"后台用户" method:"get" summary:"列表"`
	Username string `json:"username" dc:"用户名，模糊搜索"`
	Name     string `json:"name" dc:"姓名，模糊搜索"`
	Mobile   string `json:"mobile" dc:"手机号"`
	Gender   *int64 `json:"gender"  dc:"性别"`
	Status   *int64 `json:"status"  dc:"状态"`
	PageReq
}

type GetListRes struct {
	List  []*GetListItem
	Total int64
}

type GetListItem struct {
	UserId   int64  `json:"userId" dc:"用户ID"`        // 用户ID
	Username string `json:"username" dc:"用户名"`      // 用户名
	Name     string `json:"name" dc:"姓名"`            // 姓名
	Mobile   string `json:"mobile" dc:"手机号"`        // 手机号
	Gender   int64  `json:"gender" dc:"组织分类，枚举"` // 性别，枚举
	Status   int64  `json:"status" dc:"状态，枚举"`     // 状态，枚举
	CreateAt string `json:"createAt" dc:"创建时间"`
}

type GetOneReq struct {
	g.Meta `path:"/user_member" tags:"后台用户" method:"get" summary:"详情"`
	UserId int64 `json:"userId" dc:"用户ID" v:"required|integer|min:1#用户ID不能为空|用户ID必须是整数|用户ID不能小于1"`
}

type GetOneRes struct {
	UserId   int64  `json:"userId" dc:"用户ID"`         // 用户ID
	Username string `json:"username" dc:"用户名"`       // 用户名
	Name     string `json:"name" dc:"姓名"`             // 姓名
	Mobile   string `json:"mobile" dc:"手机号"`         // 手机号
	Gender   int64  `json:"gender" dc:"组织分类，枚举"`  // 性别，枚举
	Status   int64  `json:"status" dc:"状态，枚举"`      // 状态，枚举
	Super    int64  `json:"super" dc:"超级管理员，枚举"` // 超级管理员，枚举
}

type UpdatePasswordReq struct {
	g.Meta           `path:"/user_member/updatePassword" tags:"后台用户" method:"post" summary:"修改密码"`
	UserId           int64  `json:"userId" dc:"用户ID" v:"required|integer|min:1#用户ID不能为空|用户ID必须是整数|用户ID不能小于1"`
	OriginalPassword string `json:"originalPassword" dc:"原密码" v:"required#原密码不能为空"`
	NewPassword      string `json:"newPassword" dc:"新密码" v:"required|min-length:8#新密码不能为空|新密码最小长度8"`
}

type UpdatePasswordRes struct {
}
