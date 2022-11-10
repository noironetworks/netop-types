package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ErrorPodsReportingSpec defines the desired state of ErrorPodsReporting
type ErrorPodsReportingSpec struct {
	ReportingEnable bool `json:"reporting_enable"`
}

type ErrorPod struct {
	PodYaml    string `json:"pod_yaml,omitempty"`
	PodEventes string `json:"events,omitempty"`
}

// ErrorPodsReporting is the Schema for the errorpodsreportings API
type ErrorPodsReporting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ErrorPodsReportingSpec   `json:"spec,omitempty"`
	Status ErrorPodsReportingStatus `json:"status,omitempty"`
}

// ErrorPodsReportingStatus defines the observed state of ErrorPodsReporting
type ErrorPodsReportingStatus struct {
	UpdatedAt string              `json:"updated_at"`
	ErrorPods map[string]ErrorPod `json:"error_pods"`
}

//+kubebuilder:object:root=true

// ErrorPodsReportingList contains a list of ErrorPodsReporting CR
type ErrorPodsReportingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ErrorPodsReporting `json:"items"`
}
