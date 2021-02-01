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

package config

import (
	"github.com/Tencent/bk-bcs/bcs-common/common/conf"
)

// AlertManagerOptions parse command-line parameters to options
type AlertManagerOptions struct {
	conf.FileConfig
	conf.MetricConfig
	conf.ServiceConfig
	conf.CertConfig
	conf.LogConfig
	conf.ProcessConfig

	EtcdOptions        *EtcdOptions        `json:"etcdOptions"`
	AlertServerOptions *AlertServerOptions `json:"alertServerOptions"`
	ServerAddress      string              `json:"apigateway-addr" value:"" usage:"bcs apigateway address"`
	QueueConfig        string              `json:"queue_config_file" value:"queue.conf" usage:"Config file for queue."`
	DebugMode          bool                `json:"debug_mode" value:"false" usage:"Debug mode, use pprof."`

	ResourceSubs       []ResourceSubType   `json:"resourceSubs" value:"" usage:"ResourceSubs consumer"`
}

type EtcdOptions struct {
	EtcdServers  string `json:"etcd-servers" value:"127.0.0.1:2379" usage:"List of etcd servers to connect with (scheme://ip:port)"`
	EtcdCaFile   string `json:"etcd-cafile" value:"./etcd/ca.pem" usage:"SSL certificate ca file"`
	EtcdCertFile string `json:"etcd-certfile" value:"./etcd/client.pem" usage:"SSL certificate cert file"`
	EtcdKeyFile  string `json:"etcd-keyfile" value:"./etcd/client-key.pem" usage:"SSL certificate cert-key file"`
}

type AlertServerOptions struct {
	Server      string `json:"server"`
	AppCode     string `json:"appCode"`
	AppSecret   string `json:"appSecret"`
	ServerDebug bool   `json:"debugLevel"`
}

type ResourceSubType struct {
	Switch   string `json:"switch"`
	Category string `json:"category"`
}

//NewAlertManagerOptions create AlertManagerOptions object
func NewAlertManagerOptions() *AlertManagerOptions {
	return &AlertManagerOptions {}
}
