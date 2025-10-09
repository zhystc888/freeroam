package admin_member

import (
	v1 "bbk/app/user/api/admin_member/v1"
	_ "bbk/app/user/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "bbk/app/user/internal/logic"

	_ "bbk/app/user/boot"
	"context"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func newUAdminMemberControllerForTest() (*Controller, context.Context) {
	ctx := context.TODO()
	c := newAdminMemberController()
	return c, ctx
}

func TestController_Get(t *testing.T) {
	// 获取对象和上下文
	s, ctx := newUAdminMemberControllerForTest()

	// 测试开始
	gtest.C(t, func(t *gtest.T) {
		// 1. 返回正常
		res, err := s.Get(ctx, &v1.GetReq{
			UserId: 1,
		})
		t.Assert(err, nil)
		if res != nil && res.UserId != 1 {
			t.Error()
		}

		// 2. 返回空
		res, err = s.Get(ctx, &v1.GetReq{
			UserId: 0,
		})
		t.Assert(err, nil)
		t.AssertNil(res)
	})
}
