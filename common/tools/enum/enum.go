package enum

import (
	"context"
	v1 "freeroam/app/system/api/enum/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx        context.Context
	enumClient v1.EnumClient
)

func init() {
	conn := grpcx.Client.MustNewGrpcClientConn("system")

	enumClient = v1.NewEnumClient(conn)
	ctx = gctx.GetInitCtx()
}

// GetByType 获取指定一个 type 枚举
func GetByType(enumType string) (*v1.GetByTypeOptionList, error) {
	emu, err := enumClient.GetByType(ctx, &v1.GetByTypeReq{EnumTypes: []string{enumType}})
	if err != nil {
		return nil, err
	}

	return emu.GetByTypeOptionListMap[enumType], nil
}

// GetByTypes 获取指定多个 type 枚举
func GetByTypes(enumTypes []string) (map[string]*v1.GetByTypeOptionList, error) {
	emu, err := enumClient.GetByType(ctx, &v1.GetByTypeReq{EnumTypes: enumTypes})
	if err != nil {
		return nil, err
	}

	return emu.GetByTypeOptionListMap, nil
}

// GetByTypeAndCode 获取指定多个个 type 枚举
func GetByTypeAndCode(enumType, enumCode string) (*v1.GetByTypeAndCodeRes, error) {
	emu, err := enumClient.GetByTypeAndCode(ctx,
		&v1.GetByTypeAndCodeReq{
			EnumType: enumType,
			EnumCode: enumCode,
		})
	if err != nil {
		return nil, err
	}

	return emu, nil
}
