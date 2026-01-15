package packed

import (
	_ "freeroam/app/system/boot"
	_ "freeroam/app/system/internal/logic"
	_ "freeroam/common/berror"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
)
