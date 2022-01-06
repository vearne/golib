package metric

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type key string

const (
	StartTime key = "_start"
)



type DurationHook struct {
	redisCollector *RedisCollector
	role string
}


func NewDurationHook(redisCollector *RedisCollector, role string) *DurationHook{
	return &DurationHook{redisCollector:redisCollector, role:role}
}

func (hook *DurationHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	start := time.Now()
	ctx = context.WithValue(ctx, StartTime, start)
	return ctx, nil
}
func (hook *DurationHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	start := ctx.Value(StartTime).(time.Time)
	hook.redisCollector.requestDurationHistogram.WithLabelValues(hook.role).
		Observe(time.Since(start).Seconds())
	return nil
}

func (hook *DurationHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	start := time.Now()
	ctx = context.WithValue(ctx, StartTime, start)
	return ctx, nil
}

func (hook *DurationHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	start := ctx.Value(StartTime).(time.Time)
	hook.redisCollector.requestDurationHistogram.WithLabelValues(hook.role).
		Observe(time.Since(start).Seconds())
	return nil
}
