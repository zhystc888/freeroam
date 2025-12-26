package enum

import (
	"context"
	"encoding/json"
	v1 "freeroam/app/system/api/enum/v1"
	"freeroam/app/system/internal/consts"
	"freeroam/app/system/internal/dao"
	"freeroam/app/system/internal/model/entity"
	"freeroam/app/system/internal/service"
	"freeroam/common/berror"
	"sort"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sEnum struct{}

func init() {
	service.RegisterEnum(&sEnum{})
}

// GetByTypeAndCode 根据枚举类型和code 获取一个枚举
func (s *sEnum) GetByTypeAndCode(ctx context.Context, in *v1.GetByTypeAndCodeReq) (*v1.GetByTypeAndCodeRes, error) {
	redisEnum, err := s.redisGetByTypeAndCode(ctx, in.EnumType, in.EnumCode)
	if err != nil {
		return nil, err
	}

	if redisEnum == nil {
		redisEnum, err = s.dbToRedisByTypeAndCode(ctx, in.EnumType, in.EnumCode)
		if err != nil {
			return nil, err
		}

		if redisEnum == nil {
			return nil, nil
		}
	}

	return &v1.GetByTypeAndCodeRes{
		EnumValue:     redisEnum.EnumValue,
		EnumLabel:     redisEnum.EnumLabel,
		EnumValueDesc: redisEnum.EnumValueDesc,
		Sort:          int64(redisEnum.Sort),
	}, nil
}

// GetByType 根据枚举类型数组 获取多个枚举
func (s *sEnum) GetByType(ctx context.Context, in *v1.GetByTypeReq) (*v1.GetByTypeRes, error) {
	var res v1.GetByTypeRes
	enumTypeMap := make(map[string]*v1.GetByTypeOptionList, len(in.Type))
	for _, enumType := range in.Type {
		enums, err := s.redisGetByType(ctx, enumType)
		if err != nil {
			return &res, err
		}

		if enums == nil {
			enums, err = s.dbToRedisByType(ctx, enumType)
			if err != nil {
				return &res, err
			}

			if enums == nil {
				continue
			}
		}

		sort.Slice(enums, func(i, j int) bool {
			return enums[i].Sort < enums[j].Sort
		})

		enumTypeOptionList := make([]*v1.GetByTypeOption, len(enums))
		for i, enum := range enums {
			enumTypeOptionList[i] = &v1.GetByTypeOption{
				EnumCode:      enum.EnumCode,
				EnumValue:     enum.EnumValue,
				EnumLabel:     enum.EnumLabel,
				EnumValueDesc: enum.EnumValueDesc,
				Sort:          int64(enum.Sort),
			}
		}

		enumTypeMap[enumType] = &v1.GetByTypeOptionList{Options: enumTypeOptionList}
	}

	res.UserStatusMap = enumTypeMap
	return &res, nil
}

func (*sEnum) redisGetByTypeAndCode(ctx context.Context, enumType, enumCode string) (*entity.SystemEnumData, error) {
	enumItem, err := g.Redis().HGet(ctx, consts.RedisEnumKey+enumType, enumCode)
	if err != nil {
		return nil, gerror.NewCode(berror.RedisErr, err.Error())
	}

	if enumItem.IsEmpty() {
		return nil, nil
	}

	var data entity.SystemEnumData
	if err = json.Unmarshal(enumItem.Bytes(), &data); err != nil {
		return nil, gerror.NewCode(berror.JSONErr, err.Error())
	}

	return &data, nil
}

func (*sEnum) redisGetByType(ctx context.Context, enumType string) ([]*entity.SystemEnumData, error) {
	enumMap, err := g.Redis().HGetAll(ctx, consts.RedisEnumKey+enumType)
	if err != nil {
		return nil, gerror.NewCode(berror.RedisErr, err.Error())
	}

	if enumMap.IsEmpty() {
		return nil, nil
	}

	enumData := make([]*entity.SystemEnumData, len(enumMap.Map()))
	for _, item := range enumMap.Map() {
		var data entity.SystemEnumData
		gconv.Struct()
		if err = json.Unmarshal(item, &data); err != nil {
			return nil, gerror.NewCode(berror.JSONErr, err.Error())
		}
	}

	return enumData, nil
}

func (*sEnum) dbToRedisByTypeAndCode(ctx context.Context, enumType, enumCode string) (*entity.SystemEnumData, error) {
	m := dao.SystemEnumData
	query := m.Ctx(ctx).Safe(false).
		Where(m.Columns().EnumType, enumType).
		Where(m.Columns().EnumCode, enumCode).
		Where(m.Columns().IsEnabled, true)

	var data *entity.SystemEnumData
	if err := query.Scan(&data); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	if _, err := g.Redis().HSet(ctx, consts.RedisEnumKey+enumType, g.Map{enumCode: data}); err != nil {
		return nil, gerror.NewCode(berror.RedisErr, err.Error())
	}

	return data, nil
}

func (*sEnum) dbToRedisByType(ctx context.Context, enumType string) ([]*entity.SystemEnumData, error) {
	m := dao.SystemEnumData
	query := m.Ctx(ctx).Safe(false).
		Where(m.Columns().EnumType, enumType).
		Where(m.Columns().IsEnabled, true).
		OrderAsc(m.Columns().Sort)

	var data []*entity.SystemEnumData
	if err := query.Scan(&data); err != nil {
		return nil, gerror.NewCode(berror.DBErr, err.Error())
	}

	if len(data) < 1 {
		return nil, nil
	}

	_, err := g.Redis().Del(ctx, consts.RedisEnumKey+enumType)

	enumMap := make(g.Map, len(data))
	for _, item := range data {
		enumMap[item.EnumCode] = item
	}

	_, err = g.Redis().HSet(ctx, consts.RedisEnumKey+enumType, enumMap)
	if err != nil {
		return nil, gerror.NewCode(berror.RedisErr, err.Error())
	}

	return data, nil
}
