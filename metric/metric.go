package metric

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	ErrMetric = fmt.Errorf("metric: error happens in pkg metric")
)

type getterFn func(ch chan<- prometheus.Metric, desc *prometheus.Desc, typ prometheus.ValueType)

type collector struct {
	getter getterFn

	desc      *prometheus.Desc
	valueType prometheus.ValueType
}

func newCollector(
	name, help string,
	valueType prometheus.ValueType,
	labels []string,
	getter getterFn) *collector {

	desc := prometheus.NewDesc(name, help, labels, nil)
	return &collector{
		desc:      desc,
		valueType: valueType,
		getter:    getter,
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	c.getter(ch, c.desc, c.valueType)
}
