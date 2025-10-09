// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bbk/app/org/internal/model"
	"context"
)

type (
	IOrg interface {
		Get(ctx context.Context, id int64) (res *model.OrgVo, err error)
	}
)

var (
	localOrg IOrg
)

func Org() IOrg {
	if localOrg == nil {
		panic("implement not found for interface IOrg, forgot register?")
	}
	return localOrg
}

func RegisterOrg(i IOrg) {
	localOrg = i
}
