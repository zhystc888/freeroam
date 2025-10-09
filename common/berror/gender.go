package berror

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// NewBizError 创建业务异常
func NewBizError(code gcode.Code, msg string) error {
	return gerror.NewCode(code, msg)
}

// NewInternalError 创建服务器内部异常，附带堆栈信息
func NewInternalError(err error) error {
	if err == nil {
		return gerror.NewCode(CodeInternal, "服务器内部错误")
	}
	return gerror.WrapCode(CodeInternal, err, "服务器内部错误")
}
