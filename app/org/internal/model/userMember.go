package model

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type UserMemberListDto struct {
	Username string `p:"username" description:"用户名，模糊搜索"`
	Name     string `p:"name" description:"姓名，模糊搜索"`
	Mobile   string `p:"mobile" description:"手机号"`
	Gender   *int64 `p:"gender" description:"性别"`
	Status   *int64 `p:"status" description:"状态"`
	*PageReq
}

type UserMemberListVo struct {
	g.Meta   `orm:"table:free_user_member"`
	UserId   int64     `json:"userId" orm:"user_id" description:"用户ID"`   // 用户ID
	Username string    `json:"username" orm:"username" description:"用户名"` // 用户名
	Name     string    `json:"name" orm:"name" description:"姓名"`          // 姓名
	Mobile   string    `json:"mobile" orm:"mobile" description:"手机号"`     // 手机号
	Gender   int64     `json:"gender" orm:"gender" description:"组织分类，枚举"` // 性别，枚举
	Status   int64     `json:"status" orm:"status" description:"状态，枚举"`   // 状态，枚举
	CreateAt time.Time `json:"createAt" orm:"create_at" description:"创建时间"`
}

type UserMemberVo struct {
	g.Meta   `orm:"table:free_user_member"`
	UserId   int64  `json:"userId" orm:"user_id" description:"用户ID"`   // 用户ID
	Username string `json:"username" orm:"username" description:"用户名"` // 用户名
	Name     string `json:"name" orm:"name" description:"姓名"`          // 姓名
	Mobile   string `json:"mobile" orm:"mobile" description:"手机号"`     // 手机号
	Gender   int64  `json:"gender" orm:"gender" description:"组织分类，枚举"` // 性别，枚举
	Status   int64  `json:"status" orm:"status" description:"状态，枚举"`   // 状态，枚举
	Super    int64  `json:"super" orm:"super" description:"超级管理员，枚举"`  // 超级管理员，枚举
}
