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
 *
 */

package bcs_system

import (
	"context"
	"math"
	"time"

	"github.com/TencentBlueKing/bkmonitor-kits/logger"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/prompb"
	"github.com/thanos-io/thanos/pkg/component"
	"github.com/thanos-io/thanos/pkg/store/labelpb"
	"github.com/thanos-io/thanos/pkg/store/storepb"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/component/bcs"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/component/k8sclient"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/storegw/bcs_system/source"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-monitor/pkg/storegw/clientutil"
)

// Config 配置
type Config struct{}

// BCSSystemStore implements the store node API on top of the Prometheus remote read API.
type BCSSystemStore struct {
	config *Config
}

// NewBCSSystemStore
func NewBCSSystemStore(conf []byte) (*BCSSystemStore, error) {
	var config Config

	store := &BCSSystemStore{
		config: &config,
	}
	return store, nil
}

// Info 返回元数据信息
func (s *BCSSystemStore) Info(ctx context.Context, r *storepb.InfoRequest) (*storepb.InfoResponse, error) {
	// 默认配置
	lsets := []labelpb.ZLabelSet{}

	for _, m := range AvailableNodeMetrics {
		labelSets := labels.FromMap(map[string]string{"provider": "BCS_SYSTEM", "__name__": m})
		lsets = append(lsets, labelpb.ZLabelSet{Labels: labelpb.ZLabelsFromPromLabels(labelSets)})
	}

	res := &storepb.InfoResponse{
		StoreType: component.Store.ToProto(),
		MinTime:   math.MinInt64,
		MaxTime:   math.MaxInt64,
		LabelSets: lsets,
	}
	return res, nil
}

// LabelNames 返回 labels 列表
func (s *BCSSystemStore) LabelNames(ctx context.Context, r *storepb.LabelNamesRequest) (*storepb.LabelNamesResponse, error) {
	names := []string{"__name__"}
	return &storepb.LabelNamesResponse{Names: names}, nil
}

// LabelValues 返回 label values 列表
func (s *BCSSystemStore) LabelValues(ctx context.Context, r *storepb.LabelValuesRequest) (*storepb.LabelValuesResponse, error) {
	values := []string{}
	if r.Label == "__name__" {
		values = append(values, AvailableNodeMetrics...)
	}

	return &storepb.LabelValuesResponse{Values: values}, nil
}

// Series 返回时序数据
func (s *BCSSystemStore) Series(r *storepb.SeriesRequest, srv storepb.Store_SeriesServer) error {
	logger.Infow(clientutil.DumpPromQL(r), "minTime", r.MinTime, "maxTime", r.MaxTime, "step", r.QueryHints.StepMillis, r.Step)

	// 最小1分钟维度
	step := r.QueryHints.StepMillis
	if step < 60000 {
		step = 60000
	}

	// series 数据, 这里只查询最近5分钟
	if r.SkipChunks {
		r.MaxTime = time.Now().UnixMilli()
		r.MinTime = r.MaxTime - 5*60*1000
	}

	metricName, err := clientutil.GetLabelMatchValue("__name__", r.Matchers)
	if err != nil {
		return err
	}
	if metricName == "" {
		// return errors.New("metric name is required")
		return nil
	}

	clusterId, err := clientutil.GetLabelMatchValue("cluster_id", r.Matchers)
	if err != nil {
		return err
	}

	if clusterId == "" {
		return nil
		// return errors.New("cluster_id is required")
	}

	ctx := srv.Context()

	bcsConf := k8sclient.GetBCSConfByClusterId(clusterId)
	cluster, err := bcs.GetCluster(ctx, bcsConf, clusterId)
	if err != nil {
		return err
	}

	client, err := source.ClientFactory(cluster.ClusterId)
	if err != nil {
		return err
	}

	var (
		promSeriesSet []*prompb.TimeSeries
		promErr       error
	)

	switch metricName {
	case "bcs:cluster:cpu:total":
		promSeriesSet, promErr = client.GetClusterCPUTotal(ctx, cluster.ProjectId, cluster.ClusterId, time.UnixMilli(r.MinTime), time.UnixMilli(r.MaxTime), time.Millisecond*time.Duration(step))
	case "bcs:cluster:cpu:used":
		promSeriesSet, promErr = client.GetClusterCPUUsed(ctx, cluster.ProjectId, cluster.ClusterId, time.UnixMilli(r.MinTime), time.UnixMilli(r.MaxTime), time.Millisecond*time.Duration(step))
	default:
		return nil
	}

	if promErr != nil {
		return promErr
	}

	for _, promSeries := range promSeriesSet {
		series := &clientutil.TimeSeries{TimeSeries: promSeries}
		series = series.AddLabel("__name__", metricName)
		series = series.AddLabel("cluster_id", clusterId)
		series = series.RenameLabel("bk_namespace", "namespace")
		series = series.RenameLabel("bk_pod", "pod")

		s, err := series.ToThanosSeries(r.SkipChunks)
		if err != nil {
			return err
		}
		if err := srv.Send(storepb.NewSeriesResponse(s)); err != nil {
			return err
		}
	}

	return nil
}