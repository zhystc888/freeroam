package permissions

import (
	"context"
	v1 "freeroam/app/gateway/api/permissions/v1"
	genum "freeroam/app/gateway/internal/consts/enum"
	"freeroam/app/gateway/internal/service"
	uPermissions "freeroam/app/gateway/internal/utility/permissions"
	orgPermissions "freeroam/app/org/api/permissions/v1"
	"freeroam/common/berror"
	"freeroam/common/consts/enum"
	"freeroam/common/tools/jwt_claims"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
)

type sPermissions struct {
	PermissionsRpcService orgPermissions.PermissionsClient
}

func init() {
	conn := grpcx.Client.MustNewGrpcClientConn("org")
	authRpcService := orgPermissions.NewPermissionsClient(conn)
	service.RegisterPermissions(&sPermissions{authRpcService})
}

// GetFrontPermissions 获取当前登录用户前端权限
func (s *sPermissions) GetFrontPermissions(ctx context.Context, req *v1.GetFrontPermissionsReq) (*v1.GetFrontPermissionsRes, error) {
	// 缓存中获取数据
	frontPermissions, err := uPermissions.GetFrontPermissions(ctx)
	if err != nil {
		return nil, err
	}

	// 缓存中不存在
	if frontPermissions == nil {
		frontPermissions, err = s.getLoginMemberPermissions(ctx, genum.PermissionsFront)
		if err != nil {
			return nil, err
		}

		// 存储数据到缓存中
		if err = uPermissions.SaveFrontPermissions(ctx, frontPermissions); err != nil {
			logger.Errorf("保存前端权限到缓存中失败, err:%v", err)
		}
	}

	return &v1.GetFrontPermissionsRes{Permissions: frontPermissions}, nil
}

// VeryApiPermissions 验证当前登录用户是否有指定权限
func (s *sPermissions) VeryApiPermissions(ctx context.Context, permissions string) (bool, error) {
	// 缓存中获取数据
	veryPermissions, err := uPermissions.VeryApiPermissions(ctx, permissions)
	if err != nil {
		return false, err
	}

	// 缓存中不存在
	if veryPermissions == 1 {
		apiPermissions, err := s.getLoginMemberPermissions(ctx, genum.PermissionsApi)
		if err != nil {
			return false, err
		}

		// 存储数据到缓存中
		if err = uPermissions.SaveApiPermissions(ctx, apiPermissions); err != nil {
			logger.Errorf("保存 api 限到缓存中失败, err:%v", err)
		}

		veryPermissions, err = uPermissions.VeryApiPermissions(ctx, permissions)
		if err != nil {
			return false, err
		}
	}

	return veryPermissions == 2, nil
}

// 获取当前登录用户指定类型权限
// permType: 权限类型  "freeroam/app/gateway/internal/consts/enum"
func (s *sPermissions) getLoginMemberPermissions(ctx context.Context, permType string) ([]string, error) {
	memberId := jwt_claims.GetMemberId(ctx)
	if memberId == 0 {
		return nil, gerror.NewCodef(berror.CodeTokenIsEmpty,
			"获取登录用户数据失败, memberId:%d", memberId)
	}

	// 组装对应的权限类型
	var permTypes []string
	switch permType {
	case genum.PermissionsFront:
		permTypes = append(permTypes, enum.PermissionsPage)
		permTypes = append(permTypes, enum.PermissionsComponents)
	case genum.PermissionsApi:
		permTypes = append(permTypes, enum.PermissionsInterface)
	default:
		return nil, gerror.NewCodef(berror.DataNotExist, "未知的权限类型: %s", permType)
	}

	// 调用 Org 服务 获取权限
	if s.PermissionsRpcService == nil {
		return nil, berror.NewCode(berror.ServiceNotInitialized, "OrgRpcService")
	}
	getMemberPermissionsRes, err := s.PermissionsRpcService.GetMemberPermissions(ctx, &orgPermissions.GetMemberPermissionsReq{
		MemberId: int64(memberId),
		PermType: permTypes,
	})
	if err != nil {
		return nil, err
	}

	return getMemberPermissionsRes.Permissions, nil
}
