package config

import (
	"context"
	v1 "freeroam/app/system/api/config/v1"
	"freeroam/app/system/internal/consts"
	"freeroam/app/system/internal/dao"
	"freeroam/app/system/internal/model/entity"
	"freeroam/app/system/internal/service"
	"freeroam/common/berror"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sConfig struct{}

func init() {
	config := sConfig{}
	service.RegisterConfig(&config)
}

func (s *sConfig) Boot(ctx context.Context) {
	if err := s.dbToRedis(ctx); err != nil {
		glog.Error(ctx, err)
	}
}

// GetByCode 根据 code 获取一个 配置信息
func (s *sConfig) GetByCode(ctx context.Context, in *v1.GetByCodeReq) (*v1.GetByCodeRes, error) {
	configValue, err := s.redisGetByCode(ctx, in.ConfigCode)
	if err != nil {
		return nil, err
	}

	return &v1.GetByCodeRes{
		ConfigValue: *configValue,
	}, nil
}

func (*sConfig) redisGetByCode(ctx context.Context, configCode string) (*string, error) {
	enumItem, err := g.Redis().HGet(ctx, consts.RedisSystemConfigKey, configCode)
	if err != nil {
		return nil, gerror.NewCode(berror.RedisErr, err.Error())
	}

	if enumItem.IsEmpty() {
		return nil, gerror.NewCode(berror.DataNotExist, configCode)
	}

	configValue := enumItem.String()

	return &configValue, nil
}

func (s *sConfig) dbToRedis(ctx context.Context) error {
	_, err := g.Redis().Del(ctx, consts.RedisSystemConfigKey)

	m := dao.SystemConfig
	query := m.Ctx(ctx).Safe(false).
		Where(m.Columns().IsDeleted, false)

	var data []*entity.SystemConfig
	if err := query.Scan(&data); err != nil {
		return gerror.NewCode(berror.DBErr, err.Error())
	}

	if len(data) < 1 {
		return nil
	}

	configMap := make(g.Map, len(data))

	for _, item := range data {
		configMap[item.ConfigCode] = item.ConfigValue
	}

	_, err = g.Redis().HSet(ctx, consts.RedisSystemConfigKey, configMap)
	if err != nil {
		return gerror.NewCode(berror.RedisErr, err.Error())
	}

	return nil
}
