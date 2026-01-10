package berror

import "github.com/gogf/gf/v2/errors/gcode"

// 异常code 6位，前3位和httpStatus保持一致
var (
	DataNotExist        = gcode.New(400000, "数据不存在！", nil)
	IncorrectParameters = gcode.New(400001, "参数错误！", nil)
	PasswordErr         = gcode.New(401000, "密码错误！", nil)
	CodeTokenIsEmpty    = gcode.New(401001, "token不存在！", nil)
	CodeTokenInvalid    = gcode.New(401002, "token验证失败！", nil)
	CodeTokenExpired    = gcode.New(401003, "token已过期！", nil)

	RoleNotExist                    = gcode.New(404011, "角色不存在！", nil)
	RoleCodeAlreadyExists           = gcode.New(409010, "角色编码已存在！", nil)
	RoleIsBoundCannotDeleted        = gcode.New(409011, "角色仍被职务绑定，无法删除！", nil)
	BuiltInSystemRolesCannotDeleted = gcode.New(409012, "系统内置角色不允许删除！", nil)
)

var (
	CodeInternal       = gcode.New(500000, "服务器内部错误", nil)
	RedisErr           = gcode.New(500100, "Redis 查询错误", nil)
	DBErr              = gcode.New(500200, "DB 查询错误", nil)
	SerializationErr   = gcode.New(500300, "序列化错误", nil)
	DeserializationErr = gcode.New(500301, "反序列化错误", nil)
)
