package berror

import "github.com/gogf/gf/v2/errors/gcode"

var (
	DataNotExist = gcode.New(10300, "数据不存在！", nil)

	PasswordErr = gcode.New(10400, "密码错误！", nil)

	CodeTokenIsEmpty = gcode.New(10401, "token不存在！", nil)
	CodeTokenInvalid = gcode.New(10401, "token验证失败！", nil)
	CodeTokenExpired = gcode.New(10401, "token已过期！", nil)
)

var (
	CodeInternal = gcode.New(50000, "服务器内部错误", nil)

	RedisErr = gcode.New(50100, "Redis 查询错误", nil)

	DBErr = gcode.New(50200, "Redis 查询错误", nil)

	JSONErr = gcode.New(50300, "JSON 解析错误", nil)
)
