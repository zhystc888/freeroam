package middleware

import (
	"os"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

// ErrorHandler 统一异常处理：500 前缀的 code 返回泛化提示，其他返回 code 的 Message。
func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()

	err := r.GetError()
	if err == nil {
		return
	}

	code := gerror.Code(err)
	if code == nil {
		code = gcode.New(500000, "服务器内部错误", nil)
	}

	httpStatus := code.Code() / 1000
	if httpStatus < 100 || httpStatus >= 600 {
		httpStatus = 500
	}

	traceId := gctx.CtxId(r.Context())
	// 记录完整异常链
	g.Log().Errorf(r.Context(), "Error occurred: %+v", err)

	respCode := code.Code()
	respMsg := code.Message()
	if httpStatus == 500 {
		respMsg = "服务器内部错误"
	}

	r.Response.ClearBuffer()
	r.Response.WriteStatus(httpStatus, nil)
	body := g.Map{
		"code":    respCode,
		"message": respMsg,
		"traceId": traceId, // 返回 traceId
	}

	if os.Getenv("APP_ENV") != "prod" {
		body["data"] = g.Map{
			"error": err.Error(),
		}
	}
	r.Response.WriteJsonExit(body)
}
