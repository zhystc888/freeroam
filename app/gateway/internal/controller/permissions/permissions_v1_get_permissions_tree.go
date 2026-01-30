package permissions

import (
	"context"
	"freeroam/app/gateway/api/permissions/v1"
	oPermissions "freeroam/app/org/api/permissions/v1"
)

func (c *ControllerV1) GetPermissionsTree(ctx context.Context, req *v1.GetPermissionsTreeReq) (res *v1.GetPermissionsTreeRes, err error) {
	rpcReq := &oPermissions.GetPermissionsTreeReq{}

	rpcRes, err := c.PermissionsRpcService.GetPermissionsTree(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.GetPermissionsTreeRes{
		Tree: convertTreeNodes(rpcRes.Tree),
	}, nil
}

// convertTreeNodes 转换权限树节点
func convertTreeNodes(nodes []*oPermissions.PermissionsTreeNode) []*v1.PermissionsTreeNode {
	if nodes == nil {
		return make([]*v1.PermissionsTreeNode, 0)
	}

	result := make([]*v1.PermissionsTreeNode, len(nodes))
	for i, node := range nodes {
		result[i] = &v1.PermissionsTreeNode{
			Id:       node.Id,
			PermCode: node.PermCode,
			Name:     node.Name,
			PermType: node.PermType,
			Children: convertTreeNodes(node.Children),
		}
	}
	return result
}
