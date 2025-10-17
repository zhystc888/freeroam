package enum

import (
	"context"

	"freeroam/app/gateway/api/enum/v1"
	sEnum "freeroam/app/system/api/enum/v1"
)

func (c *ControllerV1) GetEnumList(ctx context.Context, req *v1.GetEnumListReq) (res *v1.GetEnumListRes, err error) {
	list := make([]*sEnum.GetEnumListsReqItem, len(req.List))
	for i, v := range req.List {
		list[i] = &sEnum.GetEnumListsReqItem{
			Code:      v.Code,
			TableName: v.TableName,
			Module:    v.Module,
		}
	}
	dto := &sEnum.GetEnumListsReq{
		List: list,
	}
	result, err := c.EnumRpcService.GetEnumLists(ctx, dto)

	if err != nil {
		return
	}

	res = &v1.GetEnumListRes{
		List: make(map[string][]*v1.GetEnumListResItem),
	}

	for _, v := range result.GetList() {
		item := make([]*v1.GetEnumListResItem, len(v.List))
		for index, resItem := range v.List {
			item[index] = &v1.GetEnumListResItem{
				Value: resItem.Value,
				Name:  resItem.Name,
			}
		}
		res.List[v.Name] = item
	}

	return
}
