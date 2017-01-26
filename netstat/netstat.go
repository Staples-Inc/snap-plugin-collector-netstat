/*
Copyright 2016 Staples, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package netstat

import (
	"fmt"
	"syscall"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"github.com/shirou/gopsutil/net"
)

const (
	vendor     = "staples"
	pluginName = "netstat"
)

// NetstatCollector type
type NetstatCollector struct{}

func (n *NetstatCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	cpolicy := plugin.NewConfigPolicy()
	return *cpolicy, nil
}

func (n *NetstatCollector) GetMetricTypes(plugin.Config) ([]plugin.Metric, error) {
	fields, err := getStats()
	if err != nil {
		return nil, fmt.Errorf("Error collecting metrics: %v", err)
	}

	metrics := make([]plugin.Metric, 0, len(fields))
	for name := range fields {
		metrics = append(metrics, plugin.Metric{
			Namespace: plugin.NewNamespace(vendor, "procfs", "netstat", name),
			Timestamp: time.Now(),
		})
	}
	return metrics, nil
}

func (n *NetstatCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	fields, err := getStats()
	if err != nil {
		return nil, fmt.Errorf("Error collecting metrics: %v", err)
	}
	metrics := make([]plugin.Metric, 0, len(mts))
	for i := range mts {
		val, err := getMapValueByNamespace(fields, mts[i].Namespace.Strings()[3:])
		if err != nil {
			return nil, fmt.Errorf("Error matching requested metrics: %v", err)
		}
		metrics = append(metrics, plugin.Metric{
			Namespace: mts[i].Namespace,
			Timestamp: time.Now(),
			Data:      val,
			Tags:      mts[i].Tags,
		})
	}
	return metrics, nil
}

//getMapValueByNamespace gets value from map by namespace given in array of strings
func getMapValueByNamespace(m map[string]interface{}, ns []string) (val interface{}, err error) {
	if len(ns) == 0 {
		return val, fmt.Errorf("Namespace length equal to zero")
	}

	current := ns[0]
	var ok bool
	if len(ns) == 1 {
		val, ok = m[current]
		if ok {
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
