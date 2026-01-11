package position

import (
	"context"
	"strconv"
	"strings"

	v1 "freeroam/app/gateway/api/position/v1"
	oPosition "freeroam/app/org/api/position/v1"
)

func (c *ControllerV1) GetPositionOptions(ctx context.Context, req *v1.GetPositionOptionsReq) (res *v1.GetPositionOptionsRes, err error) {
	// 解析逗号分隔的组织 ID
	orgIdStrs := strings.Split(req.OrgIds, ",")
	orgIds := make([]int64, 0, len(orgIdStrs))
	for _, s := range orgIdStrs {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		id, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			continue
		}
		orgIds = append(orgIds, id)
	}

	rpcReq := &oPosition.GetPositionOptionsReq{
		OrgIds:  orgIds,
		Status:  req.Status,
		Keyword: req.Keyword,
	}

	result, err := c.PositionRpcService.GetPositionOptions(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	list := make([]*v1.PositionOption, 0, len(result.List))
	for _, item := range result.List {
		list = append(list, &v1.PositionOption{
			Id:        item.Id,
			Name:      item.Name,
			Status:    item.Status,
			DataScope: item.DataScope,
		})
	}

	res = &v1.GetPositionOptionsRes{
		List: list,
	}
	return res, nil
}
