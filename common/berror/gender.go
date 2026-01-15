package berror

import (
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// NewCode 创建异常
func NewCode(code gcode.Code, text ...string) error {
	// 获取 code 的 message
	codeMessage := code.Message()

	// 如果有额外的 text，则追加到 message 中
	if len(text) > 0 {
		codeMessage = codeMessage + "：" + text[0]
	}

	// 创建新的 error，包含 code 和合并后的 message
	return gerror.NewCode(code, codeMessage)
}

// NewCodef 封装格式化参数异常
func NewCodef(code gcode.Code, format string, args ...any) error {
	// 获取 code 的 message
	codeMessage := code.Message()

	// 如果有传入 args，则格式化并拼接 codeMessage 和 format
	if len(args) > 0 {
		// 使用 fmt.Sprintf 来格式化拼接后的 message
		codeMessage = fmt.Sprintf("%s: %s", codeMessage, fmt.Sprintf(format, args...))
	} else {
		// 如果没有 args，直接拼接
		codeMessage = fmt.Sprintf("%s: %s", codeMessage, format)
	}

	// 创建并返回新的 error
	return gerror.NewCode(code, codeMessage)
}

func WrapCode(code gcode.Code, err error, text ...string) error {
	// 获取 code 的 message
	codeMessage := code.Message()

	// 如果有额外的 text，则追加到 message 中
	if len(text) > 0 {
		codeMessage = codeMessage + "：" + text[0]
	}

	if err == nil {
		return nil
	}
	return gerror.WrapCode(code, err, codeMessage)
}

func WrapCodef(code gcode.Code, err error, format string, args ...any) error {
	// 获取 code 的 message
	codeMessage := code.Message()

	// 如果有传入 args，则格式化并拼接 codeMessage 和 format
	if len(args) > 0 {
		// 使用 fmt.Sprintf 来格式化拼接后的 message
		codeMessage = fmt.Sprintf("%s: %s", codeMessage, fmt.Sprintf(format, args...))
	} else {
		// 如果没有 args，直接拼接
		codeMessage = fmt.Sprintf("%s: %s", codeMessage, format)
	}

	if err == nil {
		return nil
	}
	return gerror.WrapCode(code, err, codeMessage)
}

// NewInternalError 创建服务器内部异常，附带堆栈信息
func NewInternalError(err error) error {
	if err == nil {
		return gerror.NewCode(CodeInternal, "服务器内部错误")
	}
	return gerror.WrapCode(CodeInternal, err, "服务器内部错误")
}
