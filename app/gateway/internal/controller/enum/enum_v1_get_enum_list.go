package enum

import (
	"context"

	"freeroam/app/gateway/api/enum/v1"
	sEnum "freeroam/app/system/api/enum/v1"
)

func (c *ControllerV1) GetEnumList(ctx context.Context, req *v1.GetEnumListReq) (res *v1.GetEnumListRes, err error) {
	types := make([]string, len(req.Type))
	for i, item := range req.Type {
		types[i] = item
	}

	data, err := c.EnumRpcService.GetByType(ctx, &sEnum.GetByTypeReq{
		Type: types,
	})
	if err != nil {
		return nil, err
	}

	resMap := make(v1.GetEnumListRes, len(data.UserStatusMap))
	for key, item := range data.UserStatusMap {
		items := make([]v1.GetEnumListResItem, len(item.Options))
		for i, options := range item.Options {
			items[i] = v1.GetEnumListResItem{
				Label: options.EnumLabel,
				Value: options.EnumValue,
				Sort:  int(options.Sort),
			}
		}
		resMap[key] = items
	}

	return &resMap, nil
}
