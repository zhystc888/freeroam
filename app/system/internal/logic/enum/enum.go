package enum

import (
	"bbk/app/system/internal/dao"
	"bbk/app/system/internal/model"
	"bbk/app/system/internal/service"
	"bbk/common/berror"
	"bbk/common/tools/string"
	"context"
	"fmt"
)

type sEnum struct{}

func init() {
	service.RegisterEnum(&sEnum{})
}

func (s *sEnum) GetEnumList(ctx context.Context, dto *model.GetEnumListDto) (res *model.GetEnumListRes, err error) {

	m := dao.Enum
	query := m.Ctx(ctx).Safe(false).Where(m.Columns().TableName, dto.TableName).Where(m.Columns().Code, dto.Code)

	if dto.Module != "" {
		query.Where(m.Columns().Module, dto.Module)
	}

	strUtils := &string.StrUtils{}
	res = &model.GetEnumListRes{
		List: make([]*model.EnumListItem, 0),
		Name: fmt.Sprintf("%s%sList", dto.TableName, strUtils.CapitalizeFirst(dto.Code)),
	}

	err = query.Scan(&res.List)
	if err != nil {
		err = berror.NewInternalError(err)
	}

	return
}
