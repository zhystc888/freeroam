package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type LoginDto struct {
	Username string `json:"username" orm:"username"`
	Password string `json:"password" orm:"password"`
}

type LoginVo struct {
	Name                 string `json:"name"`
	Avatar               string `json:"avatar"`
	Token                string `json:"token"`
	ForceExpiresTimeUnix int64  `json:"forceExpiresTimeUnix"`
}

type LoginUserInfo struct {
	g.Meta   `orm:"table:bbk_user_login"`
	ID       int    `json:"id" orm:"id"`
	Password string `json:"password" orm:"password"`
	LastTime gtime.Time
	LastIp   string
}

type LoginAdminUserInfo struct {
	g.Meta   `orm:"table:bbk_admin_user"`
	UserID   int64          `json:"userId" orm:"user_id"`
	Username string         `json:"username" orm:"username"`
	Name     string         `json:"name"`
	Status   int64          `json:"status"`
	User     *LoginUserInfo `json:"user" orm:"with:id=user_id"`
}
