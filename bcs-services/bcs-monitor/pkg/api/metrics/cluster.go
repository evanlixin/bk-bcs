/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metrics

import (
	"time"

	bcsmonitor "github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/component/bcs_monitor"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/rest"
)

const (
	// PROVIDER TODO
	PROVIDER = `provider="BCS_SYSTEM"`
)

// Usage 使用量
type Usage struct {
	Used  string `json:"used"`
	Total string `json:"total"`
}

// UsageByte 使用量, bytes单位
type UsageByte struct {
	UsedByte  string `json:"used_bytes"`
	TotalByte string `json:"total_bytes"`
}

// ClusterOverviewMetric 集群概览接口
type ClusterOverviewMetric struct {
	CPUUsage    *Usage     `json:"cpu_usage"`
	DiskUsage   *UsageByte `json:"disk_usage"`
	MemoryUsage *UsageByte `json:"memory_usage"`
}

// handleClusterMetric Cluster 处理公共函数
func handleClusterMetric(c *rest.Context, promql string) (interface{}, error) {
	params := map[string]interface{}{
		"clusterId": c.ClusterId,
		"provider":  PROVIDER,
	}

	end := time.Now()
	start := end.Add(-time.Hour)
	result, err := bcsmonitor.QueryRange(c.Context, c.ProjectCode, promql, params, start, end, time.Minute)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// GetClusterOverview 集群概览数据
// @Summary  集群概览数据
// @Tags     Metrics
// @Success  200  {string}  string
// @Router   /overview [get]
func GetClusterOverview(c *rest.Context) (interface{}, error) {
	params := map[string]interface{}{
		"clusterId": c.ClusterId,
		"provider":  PROVIDER,
	}

	promqlMap := map[string]string{
		"cpu_used":     `bcs:cluster:cpu:used{cluster_id="%<clusterId>s", %<provider>s}`,
		"cpu_total":    `bcs:cluster:cpu:total{cluster_id="%<clusterId>s", %<provider>s}`,
		"memory_used":  `bcs:cluster:memory:used{cluster_id="%<clusterId>s", %<provider>s}`,
		"memory_total": `bcs:cluster:memory:total{cluster_id="%<clusterId>s", %<provider>s}`,
		"disk_used":    `bcs:cluster:disk:used{cluster_id="%<clusterId>s", %<provider>s}`,
		"disk_total":   `bcs:cluster:disk:total{cluster_id="%<clusterId>s", %<provider>s}`,
	}

	result, err := bcsmonitor.QueryMultiValues(c.Context, c.ProjectId, promqlMap, params, time.Now())
	if err != nil {
		return nil, err
	}

	m := ClusterOverviewMetric{
		CPUUsage: &Usage{
			Used:  result["cpu_used"],
			Total: result["cpu_total"],
		},
		MemoryUsage: &UsageByte{
			UsedByte:  result["memory_used"],
			TotalByte: result["memory_total"],
		},
		DiskUsage: &UsageByte{
			UsedByte:  result["disk_used"],
			TotalByte: result["disk_total"],
		},
	}

	return m, nil
}

// ClusterCPUUsage 集群 CPU 使用率
// @Summary  集群 CPU 使用率
// @Tags     Metrics
// @Success  200  {string}  string
// @Router   /cpu_usage [get]
func ClusterCPUUsage(c *rest.Context) (interface{}, error) {
	promql := `bcs:cluster:cpu:usage{cluster_id="%<clusterId>s", %<provider>s}`

	return handleClusterMetric(c, promql)

}

// ClusterMemoryUsage 集群内存使用率
// @Summary  集群内存使用率
// @Tags     Metrics
// @Success  200  {string}  string
// @Router   /memory_usage [get]
func ClusterMemoryUsage(c *rest.Context) (interface{}, error) {
	promql := `bcs:cluster:memory:usage{cluster_id="%<clusterId>s", %<provider>s}`

	return handleClusterMetric(c, promql)
}

// ClusterDiskUsage 集群磁盘使用率
// @Summary  集群磁盘使用率
// @Tags     Metrics
// @Success  200  {string}  string
// @Router   /disk_usage [get]
func ClusterDiskUsage(c *rest.Context) (interface{}, error) {
	promql := `bcs:cluster:disk:usage{cluster_id="%<clusterId>s", %<provider>s}`

	return handleClusterMetric(c, promql)
}