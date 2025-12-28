package boot

import (
	"freeroam/app/system/internal/service"

	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	// 初始化配置：将系统配置加载到 Redis
	service.Config().Boot(gctx.GetInitCtx())
}
