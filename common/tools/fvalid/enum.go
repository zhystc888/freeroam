package fvalid

import (
	"context"
	v1 "freeroam/app/system/api/enum/v1"
	"strings"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
)

func init() {
	gvalid.RegisterRule("enum", enumRule)
}

func enumRule(ctx context.Context, in gvalid.RuleFuncInput) error {
	parts := strings.SplitN(in.Rule, ":", 2)
	if len(parts) < 2 {
		return gerror.Newf(`枚举类型未指定 "%s"`, in.Rule)
	}
	enumType := parts[1]

	value := in.Value.String()

	conn := grpcx.Client.MustNewGrpcClientConn("system")
	enumClient := v1.NewEnumClient(conn)

	enumList, err := enumClient.GetByType(ctx, &v1.GetByTypeReq{
		EnumTypes: []string{enumType},
	})
	if err != nil {
		return gerror.Newf(`调用服务 "%s" 获取数据异常`, "system")
	}

	statusMap := enumList.GetByTypeOptionListMap
	list := statusMap[enumType]
	if len(list.GetOptions()) == 0 || (list.GetOptions() == nil) {
		return gerror.New("枚举类型不存在：" + enumType)
	}

	for _, item := range list.GetOptions() {
		if item.EnumValue == value {
			return nil
		}
	}

	if in.Message != "" {
		return gerror.New(in.Message)
	}
	return gerror.New("枚举值不合法")
}
