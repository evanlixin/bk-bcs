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

package v1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	// PortPoolBindingLabelKeyFormat label key prefix for port pool
	PortPoolBindingLabelKeyFormat = "portpool.%s.%s"
	// PortPoolBindingAnnotationKeyKeepDuration annotation key for keep duration of port pool binding
	PortPoolBindingAnnotationKeyKeepDuration = "keepduration.portbinding.bkbcs.tencent.com"

	// PortBindingTypeLabelKey label key for portbinding type
	PortBindingTypeLabelKey = "type.portbinding.bkbcs.tencent.com"
	// PortBindingTypeNode mark portbinding is related to node
	PortBindingTypeNode = "Node"
	// PortBindingTypePod mark portbinding is related to pod, empty PortBindingType is regarded as pod
	PortBindingTypePod = "Pod"

	// NodePortBindingConfigMapName name of node portbinding configmap, stores binding ports info of node
	NodePortBindingConfigMapName = "bcs-ingress-controller-node-port-binding"
	// NodePortBindingConfigMapNsLabel mark namespace that need injected configmap
	NodePortBindingConfigMapNsLabel = "bcs-ingress-controller-node-port-binding-configmap-inject"
	// NodePortBindingConfigMapNsLabelValue mark namespace that need injected configmap
	NodePortBindingConfigMapNsLabelValue = "true"
)

// PortBindingItem defines the port binding item
type PortBindingItem struct {
	PoolName              string                    `json:"poolName"`
	PoolNamespace         string                    `json:"poolNamespace"`
	LoadBalancerIDs       []string                  `json:"loadBalancerIDs,omitempty"`
	ListenerAttribute     *IngressListenerAttribute `json:"listenerAttribute,omitempty"`
	UptimeCheck           *UptimeCheckConfig        `json:"uptimeCheck,omitempty"`
	PoolItemLoadBalancers []*IngressLoadBalancer    `json:"poolItemLoadBalancers,omitempty"`
	PoolItemName          string                    `json:"poolItemName"`
	Protocol              string                    `json:"protocol"`
	StartPort             int                       `json:"startPort"`
	EndPort               int                       `json:"endPort"`
	RsStartPort           int                       `json:"rsStartPort"`
	// +optional
	HostPort bool   `json:"hostPort,omitempty"`
	External string `json:"external,omitempty"`
}

// GetKey get port pool item key
func (pbi *PortBindingItem) GetKey() string {
	if pbi == nil {
		return ""
	}
	return pbi.PoolItemName
}

func (pbi *PortBindingItem) GetFullKey() string {
	if pbi == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", pbi.PoolNamespace, pbi.PoolName, pbi.PoolItemName)
}

// PortBindingSpec defines the desired state of PortBinding
type PortBindingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	PortBindingList []*PortBindingItem `json:"portBindingList,omitempty"`
}

// HasEnableUptimeCheck check if has enable uptime check
func (in *PortBindingSpec) HasEnableUptimeCheck() bool {
	for _, item := range in.PortBindingList {
		if item.UptimeCheck != nil && item.UptimeCheck.IsEnabled() {
			return true
		}
	}
	return false
}

// PortBindingStatusItem port binding item status
type PortBindingStatusItem struct {
	PoolName      string `json:"portPoolName"`
	PoolNamespace string `json:"portPoolNamespace"`
	PoolItemName  string `json:"poolItemName"`
	StartPort     int    `json:"startPort"`
	EndPort       int    `json:"endPort"`
	// Status is single pool item status
	Status string `json:"status"`

	UptimeCheckStatus *UptimeCheckTaskStatus `json:"uptimeCheckStatus,omitempty"`
}

func (in *PortBindingStatusItem) GetFullKey() string {
	if in == nil {
		return ""
	}
	return fmt.Sprintf("%s/%s/%s", in.PoolNamespace, in.PoolName, in.PoolItemName)
}

// PortBindingStatus defines the observed state of PortBinding
type PortBindingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// 整体Pod绑定的状态, NotReady, PartialReady, Ready
	Status                string                   `json:"status"`
	UpdateTime            string                   `json:"updateTime"`
	PortBindingStatusList []*PortBindingStatusItem `json:"portPoolBindStatusList,omitempty"`
	PortBindingType       string                   `json:"portBindingType,omitempty"`
}

// +kubebuilder:object:root=true
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="status",type=string,JSONPath=`.status.status`

// PortBinding is the Schema for the portbindings API
type PortBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PortBindingSpec   `json:"spec,omitempty"`
	Status PortBindingStatus `json:"status,omitempty"`
}

func (pb *PortBinding) GetPortBindingType() string {
	if pb.Labels == nil {
		return PortBindingTypePod
	}

	if pType, ok := pb.Labels[PortBindingTypeLabelKey]; !ok {
		return PortBindingTypePod
	} else {
		return pType
	}
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PortBindingList contains a list of PortBinding
type PortBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PortBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PortBinding{}, &PortBindingList{})
}
