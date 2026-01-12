package berror

import "github.com/gogf/gf/v2/errors/gcode"

// 异常code 6位，前3位和httpStatus保持一致
var (
	DataNotExist = gcode.New(400000, "数据不存在！", nil)

	RoleNotExist                    = gcode.New(404011, "角色不存在！", nil)
	RoleCodeAlreadyExists           = gcode.New(409010, "角色编码已存在！", nil)
	RoleIsBoundCannotDeleted        = gcode.New(409011, "角色仍被职务绑定，无法删除！", nil)
	BuiltInSystemRolesCannotDeleted = gcode.New(409012, "系统内置角色不允许删除！", nil)

	// Token 层错误（401）
	CodeTokenIsEmpty = gcode.New(401001, "未登录（缺少 Authorization）", nil)
	CodeTokenInvalid = gcode.New(401002, "Token 验证失败", nil)

	// 登录校验失败（401｜/auth/login）
	CodeAccountNotExist    = gcode.New(400101, "账号不存在", nil)
	CodePasswordError      = gcode.New(400102, "密码错误", nil)
	CodeAccountDisabled    = gcode.New(400103, "账号被禁用", nil)
	CodeAccountResigned    = gcode.New(400104, "账号已离职", nil)
	CodeAccountDeleted     = gcode.New(400105, "账号已删除（软删 is_deleted=1）", nil)
	UnrecognizedClientType = gcode.New(400106, "未识别的客户端类型", nil)
)

// 依赖/系统异常（500）
var (
	CodeInternal           = gcode.New(500000, "服务器内部错误", nil)
	JWTSecretCannotBeEmpty = gcode.New(500001, "jwt.secret 配置项不能为空", nil)
	CodeConfigReadErr      = gcode.New(500001, "读取 JWT 配置失败", nil)
	CodeRedisErr           = gcode.New(500001, "Redis 读写失败/超时", nil)
	MissingSid             = gcode.New(500002, "缺少 sid（使用 RegisteredClaims.ID / jti）", nil)
	JWTSigningFailed       = gcode.New(500003, "JWT 签名失败", nil)
	TokenClaimsFormatErr   = gcode.New(500004, "Token Claims 格式错误或验证失败", nil)
	ServiceNotInitialized  = gcode.New(500101, "grpc服务未初始化", nil)
	RedisNotInitialized    = gcode.New(500100, "Redis 未初始化", nil)

	// 兼容旧代码（保留）
	RedisErr           = gcode.New(500100, "Redis 查询错误", nil)
	DBErr              = gcode.New(500200, "DB 查询错误", nil)
	DeserializationErr = gcode.New(500301, "反序列化错误", nil)
)
