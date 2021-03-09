package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
	"sync"
)

type MySQLCollector struct {
	Clients    map[string]*gorm.DB
	lock       sync.Mutex
	registered bool
}

var mySQLCollector *MySQLCollector

func init() {
	mySQLCollector = &MySQLCollector{
		Clients: make(map[string]*gorm.DB),
	}
}

func AddMySQL(client *gorm.DB, role string) {
	if client == nil {
		return
	}

	//client.DB().Stats()

	mySQLCollector.lock.Lock()
	defer mySQLCollector.lock.Unlock()

	if !mySQLCollector.registered {
		mySQLCollector.register()
	}

	mySQLCollector.Clients[role] = client
}

func (mc *MySQLCollector) register() {
	prometheus.MustRegister(newCollector(
		"mysql_pool_state",
		"MySQL pool state",
		prometheus.GaugeValue,
		[]string{"role", "state"},
		mc.stat,
	))

	prometheus.MustRegister(newCollector(
		"mysql_pool_fetches_total",
		"MySQL pool fetches total",
		prometheus.CounterValue,
		[]string{"role", "state"},
		mc.fetches,
	))
	mc.registered = true
}

func (mc *MySQLCollector) stat(ch chan<- prometheus.Metric, desc *prometheus.Desc, typ prometheus.ValueType) {
	for role, client := range mc.Clients {
		db, _ := client.DB()
		st := db.Stats()

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.Idle),
			role, "idle",
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.InUse),
			role, "active",
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.MaxOpenConnections),
			role, "poolsize",
		)
	}
}

func (mc *MySQLCollector) fetches(ch chan<- prometheus.Metric, desc *prometheus.Desc, typ prometheus.ValueType) {
	for role, client := range mc.Clients {
		db, _ := client.DB()
		st := db.Stats()

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.WaitCount),
			role, "wait_count",
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.WaitDuration),
			role, "wait_duration",
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.MaxIdleClosed),
			role, "max_idle_closed",
		)

		ch <- prometheus.MustNewConstMetric(
			desc,
			typ,
			float64(st.MaxLifetimeClosed),
			role, "max_life_closed",
		)
	}
}
