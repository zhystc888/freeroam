package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserBase struct {
	g.Meta   `orm:"table:bbk_user_base"`
	ID       int `json:"id" orm:"id"`
	LastTime gtime.Time
	LastIp   string
}
