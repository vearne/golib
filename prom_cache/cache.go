package prom_cache

import (
	"github.com/gogf/gf/g/os/gcache"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

// local cache
// 在local cache的基础上，增加了prometheus监控指标
type PromCache struct {
	Kind          string
	InternalCache *gcache.Cache
	PromReqTotal  *prometheus.CounterVec
	PromSize      prometheus.GaugeFunc
}

func NewCacheWithProm(kind string, maxCapacity int) *PromCache {
	c := PromCache{}
	c.Kind = kind
	c.InternalCache = gcache.New(maxCapacity)
	c.PromReqTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: PromMetricCacheRequestTotalName,
			Help: PromMetricCacheRequestTotalHelp,
			ConstLabels: map[string]string{"kind": kind},
		},
		[]string{"state"},
	)
	c.PromSize = prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name:        PromMetricCacheSizeName,
			Help:        PromMetricCacheSizeHelp,
			ConstLabels: map[string]string{"kind": kind},
		},
		func() float64 { return float64(c.InternalCache.Size()) },
	)
	prometheus.MustRegister(c.PromReqTotal, c.PromSize)
	return &c
}

func (c *PromCache) Get(key interface{}) interface{} {
	value := c.InternalCache.Get(key)
	c.PromReqTotal.With(prometheus.Labels{
		"state": All,
	}).Inc()
	if value != nil {
		c.PromReqTotal.With(prometheus.Labels{
			"state": Hit,
		}).Inc()
	}
	return value
}

func (c *PromCache) Set(key interface{}, value interface{}, duration time.Duration) {
	c.InternalCache.Set(key, value, duration)
}

func (c *PromCache) Close() {
	c.InternalCache.Close()
}

func (c *PromCache) Clear() {
	c.InternalCache.Clear()
}

func (c *PromCache) Size() int {
	return c.InternalCache.Size()
}
