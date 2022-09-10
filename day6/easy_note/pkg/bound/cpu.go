package bound

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/shirou/gopsutil/cpu"
	"net"
	"src/ByteDance/day6/easy_note/pkg/constants"
	"src/ByteDance/day6/easy_note/pkg/errno"
)

var _ remote.InboundHandler = &cpuLimitHandler{}

type cpuLimitHandler struct {
}

func (c cpuLimitHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	//TODO implement me
	return ctx, nil
}

func (c cpuLimitHandler) OnInactive(ctx context.Context, conn net.Conn) context.Context {
	//TODO implement me
	return ctx
}

func (c cpuLimitHandler) OnRead(ctx context.Context, conn net.Conn) (context.Context, error) {
	//TODO implement me
	p := cpuPercent()
	klog.CtxInfof(ctx, "current cpu is %.2g", p)
	if p > constants.CPURateLimit {
		return ctx, errno.ServiceErr.WithMessage(fmt.Sprintf("cpu = %.2g", c))
	}

	return ctx, nil
}

func (c cpuLimitHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	//TODO implement me
	return ctx, nil
}

func cpuPercent() float64 {
	percent, _ := cpu.Percent(0, false)
	return percent[0]
}
