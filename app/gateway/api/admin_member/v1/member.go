package v1

type MemberBase struct {
	UserID int64  `json:"userId" dc:"用户ID" orm:"user_id"`
	Name   string `json:"name" dc:"姓名"`
	Status int32  `json:"status" dc:"状态，见枚举" enum:"[0,1,2]"`
}
