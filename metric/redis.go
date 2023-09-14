package metric

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	AddHook(hook redis.Hook)
	PoolStats() *redis.PoolStats
}

type RedisCollector struct {
	requestDurationHistogram *prometheus.HistogramVec
	Clients                  map[string]RedisClient

	lock sync.Mutex

	requestDurationRegistered bool
	registered                bool
}

var redisCollector *RedisCollector

func init() {
	redisCollector = &RedisCollector{
		Clients: make(map[string]RedisClient),
	}
}

// metric.AddRedis(redisClient, "car", metric.RequestDuration)
func AddRedis(client RedisClient, role string) {
	if client == nil {
		return
	}

	redisCollector.lock.Lock()
	defer redisCollector.lock.Unlock()

	if !redisCollector.registered {
		redisCollector.register()
	}

	addRequestDuration(client, role)

	redisCollector.Clients[role] = client
}

func addRequestDuration(client RedisClient, role string) {
	if redisCollector.requestDurationHistogram == nil {
		redisCollector.requestDurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "redis_request_duration_seconds",
			Help:    "Redis request duration in seconds",
			Buckets: []float64{1e-03, 2.5e-03, 5e-03, 10e-03, 25e-03, 50e-03, 100e-03, 250e-03, 500e-03, 1, 2.5, 5, 10},
		}, []string{"role"})
	}
	if !redisCollector.requestDurationRegistered {
		prometheus.MustRegister(redisCollector.requestDurationHistogram)
		redisCollector.requestDurationRegistered = true
	}

	client.AddHook(NewDurationHook(redisCollector, role))
}

func (rc *RedisCollector) stat(ch chan<- prometheus.Metric, desc *prometheus.Desc, typ prometheus.ValueType) {
	for role, client := range rc.Clients {
		st := client.PoolStats()

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.IdleConns),
			role, "idle",
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.TotalConns-st.IdleConns),
			role, "active",
		)

		var poolSize float64
		switch c := client.(type) {
		case *redis.Client:
			poolSize = float64(c.Options().PoolSize)
		case *redis.ClusterClient:
			poolSize = float64(c.Options().PoolSize)

		}
		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			poolSize,
			role, "poolsize",
		)
	}
}

func (rc *RedisCollector) fetches(ch chan<- prometheus.Metric, desc *prometheus.Desc, typ prometheus.ValueType) {
	for role, client := range rc.Clients {
		st := client.PoolStats()

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.Hits),
			role, "hit",
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.Misses),
			role, "miss",
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.Timeouts),
			role, "timeout",
		)
	}
}

func (rc *RedisCollector) register() {
	prometheus.MustRegister(newCollector(
		"redis_pool_state",
		"Redis pool state",
		prometheus.GaugeValue,
		[]string{"role", "state"},
		rc.stat,
	))

	prometheus.MustRegister(newCollector(
		"redis_pool_fetches_total",
		"Redis pool fetches total",
		prometheus.CounterValue,
		[]string{"role", "state"},
		rc.fetches,
	))
	rc.registered = true
}
