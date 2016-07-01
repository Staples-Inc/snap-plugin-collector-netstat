package netstat

import (
	// log "github.com/Sirupsen/logrus"

	"fmt"
	// "github.com/intelsdi-x/snap-plugin-utilities/ns"
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"github.com/shirou/gopsutil/net"
	"syscall"
	"time"
)

const (
	vendor     = "staples"
	fs         = "netstat"
	pluginName = "netstat"
	version    = 1
	pluginType = plugin.CollectorPluginType
)

type netstatPlugin struct {
}

func New() *netstatPlugin {
	netstat := &netstatPlugin{}
	return netstat
}

//Meta returns meta data for plugin
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		pluginName,
		version,
		pluginType,
		[]string{},
		[]string{plugin.SnapGOBContentType},
		plugin.ConcurrencyCount(1))
}

func (netstat *netstatPlugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	return cpolicy.New(), nil
}

func (netstat *netstatPlugin) GetMetricTypes(cfg plugin.ConfigType) (metrics []plugin.MetricType, err error) {
	fields, err := getStats()
	if err != nil {
		return nil, fmt.Errorf("Errpr collecting metrics: %v", err)
	}

	for name, _ := range fields {
		ns := core.NewNamespace(vendor, fs, name)
		metric := plugin.MetricType{
			Namespace_: ns,
			Data_:      nil,
			Timestamp_: time.Now(),
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

func (netstat *netstatPlugin) CollectMetrics(metricTypes []plugin.MetricType) (metrics []plugin.MetricType, err error) {
	fields, err := getStats()
	if err != nil {
		return nil, fmt.Errorf("Errpr collecting metrics: %v", err)
	}

	for _, metricType := range metricTypes {
		ns := metricType.Namespace()

		val, err := getMapValueByNamespace(fields, ns[2:].Strings())
		if err != nil {
			return nil, fmt.Errorf("Errpr collecting metrics: %v", err)
		}

		metric := plugin.MetricType{
			Namespace_: ns,
			Data_:      val,
			Timestamp_: time.Now(),
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

//getMapValueByNamespace gets value from map by namespace given in array of strings
func getMapValueByNamespace(m map[string]interface{}, ns []string) (val interface{}, err error) {
	if len(ns) == 0 {
		return val, fmt.Errorf("Namespace length equal to zero")
	}

	current := ns[0]

	if len(ns) == 1 {
		if val, ok := m[current]; ok {
			return val, err
		}
		return val, fmt.Errorf("Key does not exist in map {key %s}", current)
	}

	if v, ok := m[current].(map[string]interface{}); ok {
		val, err = getMapValueByNamespace(v, ns[1:])
		return val, err
	}
	return val, err
}

func getStats() (map[string]interface{}, error) {
	conns, err := net.Connections("all")
	if err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	counts["UDP"] = 0
	for _, conn := range conns {
		if conn.Type == syscall.SOCK_DGRAM {
			counts["UDP"]++
			continue
		}
		c, ok := counts[conn.Status]
		if !ok {
			counts[conn.Status] = 0
		}
		counts[conn.Status] = c + 1
	}

	fields := map[string]interface{}{
		"tcp_established": counts["ESTABLISHED"],
		"tcp_syn_sent":    counts["SYN_SENT"],
		"tcp_syn_recv":    counts["SYN_RECV"],
		"tcp_fin_wait1":   counts["FIN_WAIT1"],
		"tcp_fin_wait2":   counts["FIN_WAIT2"],
		"tcp_time_wait":   counts["TIME_WAIT"],
		"tcp_close":       counts["CLOSE"],
		"tcp_close_wait":  counts["CLOSE_WAIT"],
		"tcp_last_ack":    counts["LAST_ACK"],
		"tcp_listen":      counts["LISTEN"],
		"tcp_closing":     counts["CLOSING"],
		"tcp_none":        counts["NONE"],
		"udp_socket":      counts["UDP"],
	}

	return fields, nil
}
