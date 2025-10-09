package model

type AdminBase struct {
	UserID int64  `json:"userId" orm:"user_id"`
	Name   string `json:"name"`
	Status int32  `json:"status"`
}
