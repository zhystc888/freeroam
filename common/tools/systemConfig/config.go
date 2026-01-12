package systemConfig

import (
	"context"
	v1 "freeroam/app/system/api/config/v1"
	"strconv"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx         context.Context
	systemClint v1.ConfigClient
)

func init() {
	conn := grpcx.Client.MustNewGrpcClientConn("system")

	systemClint = v1.NewConfigClient(conn)
	ctx = gctx.GetInitCtx()
}

// GetString 获取 string 类型配置
func GetString(code string) (*string, error) {
	config, err := systemClint.GetByCode(ctx, &v1.GetByCodeReq{ConfigCode: code})
	if err != nil {
		return nil, err
	}

	return &config.ConfigValue, nil
}

// GetBool 获取 bool 类型配置
func GetBool(code string) (*bool, error) {
	stringValue, err := GetString(code)
	if err != nil {
		return nil, err
	}

	parseBool, err := strconv.ParseBool(*stringValue)
	return &parseBool, err
}

// GetBoolD 获取 bool 类型配置（带默认值，不抛异常，不返回指针，方便调用方不再判断，慎用）
func GetBoolD(code string, defaultValue bool) bool {
	stringValue, err := GetString(code)
	if err != nil {
		return defaultValue
	}

	parseBool, err := strconv.ParseBool(*stringValue)
	if err != nil {
		return defaultValue
	}

	return parseBool
}

// GetInt 获取 int64 类型配置
func GetInt(code string) (*int64, error) {
	stringValue, err := GetString(code)
	if err != nil {
		return nil, err
	}

	i, err := strconv.ParseInt(*stringValue, 10, 64)
	return &i, err
}

// GetIntD 获取 int64 类型配置（带默认值，不可为0，不抛异常，不返回指针，方便调用方不再判断，慎用）
func GetIntD(code string, defaultValue int64) int64 {
	stringValue, err := GetString(code)
	if err != nil {
		return defaultValue
	}

	i, err := strconv.ParseInt(*stringValue, 10, 64)
	if i == 0 || err != nil {
		return defaultValue
	}

	return i
}

// GetUint 获取 uint64 类型配置
func GetUint(code string) (*uint64, error) {
	stringValue, err := GetString(code)
	if err != nil {
		return nil, err
	}

	i, err := strconv.ParseUint(*stringValue, 10, 64)
	return &i, err
}

// GetFloat 获取 float64 类型配置
func GetFloat(code string) (*float64, error) {
	stringValue, err := GetString(code)
	if err != nil {
		return nil, err
	}

	float, err := strconv.ParseFloat(*stringValue, 64)
	return &float, err
}
