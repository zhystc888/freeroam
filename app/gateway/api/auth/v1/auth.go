package v1

import "github.com/gogf/gf/v2/frame/g"

// LoginReq 登录请求
type LoginReq struct {
	g.Meta   `path:"/auth/login" tags:"认证管理" method:"post" summary:"登录"`
	Username string `json:"username" v:"required|length:1,20#账号不能为空|账号长度必须在1-20之间" dc:"账号"`
	Password string `json:"password" v:"required|length:1,255#密码不能为空|密码长度不能超过255" dc:"密码"`
}

// LoginRes 登录响应
type LoginRes struct {
	AccessToken string `json:"access_token" dc:"访问令牌(JWT)"`
	ExpireAt    int64  `json:"expire_at" dc:"绝对过期时间点(时间戳)"`
}

// LogoutReq 登出请求
type LogoutReq struct {
	g.Meta `path:"/auth/logout" tags:"认证管理" method:"post" summary:"登出"`
}

// LogoutRes 登出响应
type LogoutRes struct {
	Success bool `json:"success" dc:"是否成功"`
}
