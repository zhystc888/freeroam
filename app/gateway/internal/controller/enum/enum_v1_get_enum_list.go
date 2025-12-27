package enum

import (
	"context"
	"fmt"

	"freeroam/app/gateway/api/enum/v1"
	sEnum "freeroam/app/system/api/enum/v1"
)

func (c *ControllerV1) GetEnumList(ctx context.Context, req *v1.GetEnumListReq) (res *v1.GetEnumListRes, err error) {
	types := make([]string, 0, len(req.Type))
	for _, item := range req.Type {
		types = append(types, item)
	}

	data1, err1 := c.EnumRpcService.GetByTypeAndCode(ctx, &sEnum.GetByTypeAndCodeReq{
		EnumType: "test",
		EnumCode: "b",
	})
	fmt.Println(data1, err1)

	data, err := c.EnumRpcService.GetByType(ctx, &sEnum.GetByTypeReq{
		EnumTypes: types,
	})
	if err != nil {
		return nil, err
	}

	resMap := make(v1.GetEnumListRes, len(data.GetByTypeOptionListMap))
	for key, item := range data.GetByTypeOptionListMap {
		items := make([]v1.GetEnumListResItem, 0, len(item.Options))
		for _, options := range item.Options {
			items = append(items, v1.GetEnumListResItem{
				Label: options.EnumLabel,
				Value: options.EnumValue,
				Sort:  int(options.Sort),
			})
		}
		resMap[key] = items
	}

	return &resMap, nil
}
