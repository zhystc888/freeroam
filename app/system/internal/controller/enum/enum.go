package enum

import (
	"context"
	v1 "freeroam/app/system/api/enum/v1"
	"freeroam/app/system/internal/model"
	"freeroam/app/system/internal/service"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
)

type Controller struct {
	v1.UnimplementedEnumServer
	EnumServer service.IEnum
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterEnumServer(s.Server, &Controller{
		EnumServer: service.Enum(),
	})
}

func (c *Controller) GetEnumLists(ctx context.Context, req *v1.GetEnumListsReq) (res *v1.EnumListsRes, err error) {

	res = &v1.EnumListsRes{
		List: make([]*v1.EnumListsItem, len(req.List)),
	}

	var result *model.GetEnumListRes

	for i, v := range req.GetList() {
		dto := &model.GetEnumListDto{
			Code:      v.Code,
			TableName: v.TableName,
			Module:    v.Module,
		}

		//internal\controller\enum\enum.go:40:4: result parameter err not in scope at return
		//internal\controller\enum\enum.go:37:11: inner declaration of var err error
		result, err = c.EnumServer.GetEnumList(ctx, dto)

		if err != nil {
			return
		}

		list := make([]*v1.EnumListItem, len(result.List))
		for index, item := range result.List {
			list[index] = &v1.EnumListItem{
				Value: item.Id,
				Name:  item.Name,
			}
		}

		res.List[i] = &v1.EnumListsItem{
			Name: result.Name,
			List: list,
		}
	}

	return
}
