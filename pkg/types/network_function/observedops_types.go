/*
Copyright 2022.

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

package network_function

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ObservedOpsSpec defines the desired state of ObservedOps.
type ObservedOpsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CniMountPath   string   `json:"cniMountPath,omitempty"`
	CniConfDirs    []string `json:"cniConfPath,omitempty"`
	ForceReconcile int      `json:"forceReconcile,omitempty"`
}

// ObservedOpsStatus defines the observed state of ObservedOps.
type ObservedOpsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CniType    string `json:"cniType,omitempty"`
	CniVersion string `json:"cniVersion,omitempty"`
	Multus     bool   `json:"Multus,omitempty"`
}

const (
	Cilium   = "cilium"
	NotFound = "not-found"
	Calico   = "calico"
	Aci      = "aci-cni"
)

// ObservedOps is the Schema for the observedops API.
type ObservedOps struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ObservedOpsSpec   `json:"spec,omitempty"`
	Status ObservedOpsStatus `json:"status,omitempty"`
}

// ObservedOpsList contains a list of ObservedOps.
type ObservedOpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObservedOps `json:"items"`
}
