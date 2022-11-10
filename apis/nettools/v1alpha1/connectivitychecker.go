package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConnectivityCheckerSpec defines the desired state of ConnectivityChecker
type ConnectivityCheckerSpec struct {
	ReachabilityTestEnable ReachabilityTestEnableSpec `json:"reachability_test_enable"`
	ReachabilityTests      ReachabilityTestsSpec      `json:"reachability_tests"`
}

type ReachabilityTestEnableSpec struct {
	PeriodicSync PeriodicSyncSpec `json:"periodic_sync"`
	ExternalUrl  ExternalUrl      `json:"external_url"`
	Proxy        Proxy            `json:"proxy"`
}

type ReachabilityTestsSpec struct {
	PodToPod        PodToPodSpec        `json:"pod_to_pod"`
	PodToNode       PodToNodeSpec       `json:"pod_to_node"`
	PodToClusterIp  PodToClusterIpSpec  `json:"pod_to_clusterIP"`
	PodToService    PodToServiceSpec    `json:"pod_to_service"`
	PodToLbVip      PodToLbVipSpec      `json:"pod_to_LB_VIP"`
	PodToNodePort   PodToNodePortSpec   `json:"pod_to_NodePort"`
	PodToExternal   PodToExternalSpec   `json:"pod_to_External"`
	NodeToNode      NodeToNodeSpec      `json:"node_to_node"`
	NodeToPod       NodeToPodSpec       `json:"node_to_pod"`
	NodeToClusterIp NodeToClusterIpSpec `json:"node_to_clusterIP"`
	NodeToLbVip     NodeToLbVipSpec     `json:"node_to_LB_VIP"`
	NodeToNodePort  NodeToNodePortSpec  `json:"node_to_NodePort"`
	NodeToExternal  NodeToExternalSpec  `json:"node_to_External"`
}

type PeriodicSyncSpec struct {
	// +kubebuilder:default=true
	EnableUpdates bool `json:"enable_updates"`

	// +kubebuilder:default=120
	Interval int `json:"interval"`
}

type ExternalUrl struct {
	// +kubebuilder:default="google.com"
	Url string `json:"url"`
}

type Proxy struct {
	Http_proxy  string `json:"http_proxy"`
	Https_proxy string `json:"https_proxy"`
}

type PodToPodSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type PodToNodeSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type PodToClusterIpSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type PodToServiceSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type PodToLbVipSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type PodToNodePortSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type PodToExternalSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type NodeToNodeSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type NodeToPodSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type NodeToClusterIpSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type NodeToLbVipSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type NodeToNodePortSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

type NodeToExternalSpec struct {
	// +kubebuilder:default=true
	UpdateStatus bool `json:"update_status"`
}

// ConnectivityCheckerStatus defines the observed state of ConnectivityChecker
type ConnectivityCheckerStatus struct {
	UpdatedAt   string                         `json:"updated_at"`
	Results     ConnectivityCheckerResults     `json:"results"`
	ResultCount ConnectivityCheckerResultCount `json:"result_count"`
}

type ConnectivityCheckerResultCount struct {
	Passed  int `json:"passed"`
	Failed  int `json:"failed"`
	Skipped int `json:"skipped"`
}

type ConnectivityCheckerResults struct {
	PodToPod        string `json:"pod_to_pod"`
	PodToNode       string `json:"pod_to_node"`
	PodToClusterIp  string `json:"pod_to_clusterIP"`
	PodToService    string `json:"pod_to_service"`
	PodToLbVip      string `json:"pod_to_LB_VIP"`
	PodToNodePort   string `json:"pod_to_NodePort"`
	PodToExternal   string `json:"pod_to_External"`
	NodeToNode      string `json:"node_to_node"`
	NodeToPod       string `json:"node_to_pod"`
	NodeToClusterIp string `json:"node_to_clusterIP"`
	NodeToLbVip     string `json:"node_to_LB_VIP"`
	NodeToNodePort  string `json:"node_to_NodePort"`
	NodeToExternal  string `json:"node_to_External"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConnectivityChecker is the Schema for the connectivitycheckers API
type ConnectivityChecker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectivityCheckerSpec   `json:"spec,omitempty"`
	Status ConnectivityCheckerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConnectivityCheckerList contains a list of ConnectivityChecker
type ConnectivityCheckerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectivityChecker `json:"items"`
}
