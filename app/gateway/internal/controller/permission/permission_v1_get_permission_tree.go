package permission

import (
	"context"

	v1 "freeroam/app/gateway/api/permission/v1"
	oPermission "freeroam/app/org/api/permission/v1"
)

func (c *ControllerV1) GetPermissionTree(ctx context.Context, req *v1.GetPermissionTreeReq) (res *v1.GetPermissionTreeRes, err error) {
	rpcReq := &oPermission.GetPermissionTreeReq{}

	rpcRes, err := c.PermissionRpcService.GetPermissionTree(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.GetPermissionTreeRes{
		Tree: convertTreeNodes(rpcRes.Tree),
	}, nil
}

// convertTreeNodes 转换权限树节点
func convertTreeNodes(nodes []*oPermission.PermissionTreeNode) []*v1.PermissionTreeNode {
	if nodes == nil {
		return nil
	}

	result := make([]*v1.PermissionTreeNode, len(nodes))
	for i, node := range nodes {
		result[i] = &v1.PermissionTreeNode{
			Id:       node.Id,
			PermCode: node.PermCode,
			Name:     node.Name,
			PermType: node.PermType,
			Children: convertTreeNodes(node.Children),
		}
	}
	return result
}
