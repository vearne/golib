package metric

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net"
	"time"
)

type DurationHook struct {
	redisCollector *RedisCollector
	role           string
}

func NewDurationHook(redisCollector *RedisCollector, role string) *DurationHook {
	return &DurationHook{redisCollector: redisCollector, role: role}
}

func (hook *DurationHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

func (hook *DurationHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		start := time.Now()
		defer hook.redisCollector.requestDurationHistogram.WithLabelValues(hook.role).
			Observe(time.Since(start).Seconds())
		return next(ctx, cmd)
	}
}

func (hook *DurationHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		start := time.Now()
		defer hook.redisCollector.requestDurationHistogram.WithLabelValues(hook.role).
			Observe(time.Since(start).Seconds())
		return next(ctx, cmds)
	}
}
