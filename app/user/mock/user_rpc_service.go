package mock

import (
	uam "bbk/app/user/api/admin_member/v1"
	"bbk/common/berror"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"google.golang.org/grpc"
)

type UserRpcServiceMock struct{}

func (m *UserRpcServiceMock) Get(ctx context.Context, req *uam.GetReq, opts ...grpc.CallOption) (*uam.GetRes, error) {
	val := gconv.String(ctx.Value("value"))
	fmt.Println("ctxval", val)
	// 1. 空数据
	if val == "empty" {
		return nil, nil
	}

	// 2. 异常返回
	if val == "error" {
		return nil, berror.NewInternalError(nil)
	}

	// 3. 正常返回
	return &uam.GetRes{
		UserId:   req.UserId,
		Username: fmt.Sprintf("username%d", req.UserId),
		Name:     fmt.Sprintf("用户%d", req.UserId),
		Status:   1,
	}, nil
}

func (m *UserRpcServiceMock) GetList(ctx context.Context, req *uam.GetListReq, opts ...grpc.CallOption) (*uam.GetListRes, error) {
	val := gconv.String(ctx.Value("value"))
	// 1. 空数据
	if val == "empty" {
		return &uam.GetListRes{}, nil
	}

	// 2. 异常返回
	if val == "error" {
		return nil, berror.NewInternalError(nil)
	}

	// 3. 正常返回
	return &uam.GetListRes{
		List: []*uam.GetListItem{
			{
				UserId:   req.UserId,
				Name:     req.Name,
				Status:   1,
				LastTime: "2025-07-03 00:00:00",
				LastIp:   "127.0.0.1",
				CreateAt: "2025-06-03 00:00:00",
			},
		},
		Total: 100,
	}, nil
}
