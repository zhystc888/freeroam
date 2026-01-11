package org

import (
	"context"

	v1 "freeroam/app/gateway/api/org/v1"
	oOrg "freeroam/app/org/api/org/v1"
)

func (c *ControllerV1) GetOrgTree(ctx context.Context, req *v1.GetOrgTreeReq) (res *v1.GetOrgTreeRes, err error) {
	rpcReq := &oOrg.GetOrgTreeReq{
		Keyword: req.Keyword,
	}

	result, err := c.OrgRpcService.GetOrgTree(ctx, rpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GetOrgTreeRes{
		List: convertOrgTreeNodes(result.List),
	}
	return res, nil
}

// convertOrgTreeNodes 转换组织树节点
func convertOrgTreeNodes(nodes []*oOrg.OrgTreeNode) []*v1.OrgTreeNode {
	if nodes == nil {
		return nil
	}
	result := make([]*v1.OrgTreeNode, 0, len(nodes))
	for _, node := range nodes {
		result = append(result, &v1.OrgTreeNode{
			Id:          node.Id,
			Name:        node.Name,
			FullName:    node.FullName,
			Code:        node.Code,
			Category:    node.Category,
			Status:      node.Status,
			MemberCount: node.MemberCount,
			Children:    convertOrgTreeNodes(node.Children),
		})
	}
	return result
}
