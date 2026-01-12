package packed

import (
	_ "freeroam/app/gateway/boot"
	_ "freeroam/common/tools/fvalid"

	_ "freeroam/app/gateway/internal/logic"

	_ "freeroam/common/berror"

	// Redis driver (g.Redis()).
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)
