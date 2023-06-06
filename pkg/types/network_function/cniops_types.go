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

// CniOpsSpec defines the desired state of CniOps.
type CniOpsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Config       Config              `json:"config,omitempty"`
	ArrayOptions map[string][]string `json:"arrayOptions,omitempty"`
}

// CniOpsStatus defines the observed state of CniOps.
type CniOpsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CniType            string            `json:"cniType,omitempty"`
	CniVersion         string            `json:"cniVersion,omitempty"`
	CniStatus          map[string]string `json:"cniStatus,omitempty"`
	State              State             `json:"state,omitempty"`
	ManagedState       CniManagedState   `json:"managedState,omitempty"`
	UpgradeStatus      CniUpgradeStatus  `json:"upgradeStatus,omitempty"`
	CniState           string            `json:"cniState,omitempty"`
	ObservedGeneration int64             `json:"observedGeneration,omitempty"`
	WorkloadCheck      map[string]string `json:"workloadCheck,omitempty"`
	Ipam               string            `json:"ipam,omitempty"`
}

type CniManagedState string

const (
	New       CniManagedState = "New"
	Unmanaged CniManagedState = "Unmanaged" // Observed
	Imported  CniManagedState = "Imported"
)

type CniUpgradeStatus struct {
	CniUpgradeState CniUpgradeState `json:"cniUpgradeState,omitempty"`
	PreviousVersion string          `json:"previousVersion,omitempty"`
	CurrentVersion  string          `json:"currentVersion,omitempty"`
}

type CniUpgradeState string

const (
	NoUpgrade CniUpgradeState = "None"
	Upgrade   CniUpgradeState = "Upgrade"
	Downgrade CniUpgradeState = "Downgrade"
)

// CniOps is the Schema for the cniops API
type CniOps struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CniOpsSpec   `json:"spec,omitempty"`
	Status CniOpsStatus `json:"status,omitempty"`
}

// CniOpsList contains a list of CniOps.
type CniOpsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CniOps `json:"items"`
}
