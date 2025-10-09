package admin_member

import (
	"bbk/app/gateway/api/admin_member"
	v1 "bbk/app/gateway/api/admin_member/v1"
	"bbk/app/user/mock"
	"context"
	"github.com/gogf/gf/v2/test/gtest"
	"testing"
)

func newUserControllerForTest() (admin_member.IAdminMemberV1, context.Context) {
	ctx := context.TODO()
	c := &ControllerV1{
		&mock.UserRpcServiceMock{},
	}
	return c, ctx
}

func TestControllerV1_Get(t *testing.T) {
	// 获取对象和上下文
	s, ctx := newUserControllerForTest()
	// 测试开始
	gtest.C(t, func(t *gtest.T) {
		// 1. 正常返回
		ctx = context.WithValue(ctx, "value", "")
		res, err := s.Get(ctx, &v1.GetReq{
			UserID: 1,
		})
		t.AssertNil(err)
		t.Assert(res.UserID, 1)

		// 2. 异常返回
		ctx = context.WithValue(ctx, "value", "error")
		_, err = s.Get(ctx, &v1.GetReq{
			UserID: 1,
		})
		t.AssertNE(err, nil)

		// 3. 空数据返回
		ctx = context.WithValue(ctx, "value", "empty")
		res, err = s.Get(ctx, &v1.GetReq{
			UserID: 1200,
		})
		t.AssertNil(err)
		t.AssertNil(res)
	})
}

func TestControllerV1_GetList(t *testing.T) {
	// 获取对象和上下文
	s, ctx := newUserControllerForTest()
	// 测试开始
	gtest.C(t, func(t *gtest.T) {
		// 1. 正常返回
		ctx = context.WithValue(ctx, "value", "")
		res, err := s.GetMemberList(ctx, &v1.GetMemberListReq{})
		t.AssertNil(err)
		t.AssertGT(len(res.List), 0)
		t.AssertGT(res.Total, 0)

		// 2. 异常返回
		ctx = context.WithValue(ctx, "value", "error")
		_, err = s.GetMemberList(ctx, &v1.GetMemberListReq{})
		t.AssertNE(err, nil)

		// 3. 空数据返回
		ctx = context.WithValue(ctx, "value", "empty")
		res, err = s.GetMemberList(ctx, &v1.GetMemberListReq{})
		t.AssertNil(err)
		t.Assert(*res, v1.GetMemberListRes{List: make([]v1.GetMemberListItem, 0), Total: 0})
	})
}
