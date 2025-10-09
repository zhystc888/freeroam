package org

import (
	"bbk/app/org/internal/dao"
	"bbk/app/org/internal/model"
	"bbk/app/org/internal/service"
	"bbk/common/berror"
	"context"
	"fmt"
)

type sOrg struct{}

func init() {
	service.RegisterOrg(&sOrg{})
}

func (s *sOrg) Get(ctx context.Context, id int64) (res *model.OrgVo, err error) {
	m := dao.OrgStructure
	err = m.Ctx(ctx).WithAll().Where(m.Columns().Id, id).Scan(&res)
	fmt.Println(res.Supervisor)
	// 封装error为制定异常信息
	if err != nil {
		err = berror.NewInternalError(err)
	}
	return
}
