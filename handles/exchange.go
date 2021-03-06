package handles

import (
	"github.com/henrylee2cn/opay"
)

/*
 * 兑换
 */
type Exchange struct {
	Background
}

// 编译期检查接口实现
var _ Handler = (*Exchange)(nil)

// 执行入口
func (e *Exchange) ServeOpay(ctx *opay.Context) error {
	if !ctx.HasStakeholder() {
		return opay.ErrStakeholderNotExist
	}
	if ctx.GreaterOrEqual(ctx.Request.Initiator.GetAmount(), 0) ||
		ctx.SmallerOrEqual(ctx.Request.Stakeholder.GetAmount(), 0) {
		return opay.ErrIncorrectAmount
	}
	return e.Call(e, ctx)
}

// 处理账户并标记订单为成功状态
func (e *Exchange) Succeed() error {
	// 操作账户
	err := e.Background.Context.UpdateBalance()
	if err != nil {
		return err
	}

	// 更新订单
	return e.Background.Context.Succeed()
}

// 实时兑换
func (e *Exchange) SyncDeal() error {
	// 操作账户
	err := e.Background.Context.UpdateBalance()
	if err != nil {
		return err
	}

	return e.Background.Context.SyncDeal()
}
