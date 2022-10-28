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

package controller

import (
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InstallerSpec defines the desired state of Installer.
type InstallerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	OperatorConfigs OperatorConfigs `json:"operatorConfigs,omitempty"`
}

type OperatorConfigs struct {
	Platform   map[PlatformOperators]OperatorInfo   `json:"platform,omitempty"`
	Networking map[NetworkingOperators]OperatorInfo `json:"networking,omitempty"`
}

type OperatorInfo struct {
	OPSubType        OPSubType           `json:"type,omitempty"`
	Options          map[string]string   `json:"options,omitempty"`
	ArrayOptions     map[string][]string `json:"arrayOptions,omitempty"`
	WorkloadVersion  string              `json:"version,omitempty"`
	ManagedComponent bool                `json:"managedComponent,omitempty"`
}

type PlatformOperators string

const (
	Observed   PlatformOperators = "cko-observed"
	Diagnostic PlatformOperators = "cko-diagnostic"
	G          PlatformOperators = "cko-gitops"
)

func (p PlatformOperators) IsValid() error {
	switch p {
	case Observed, G, Diagnostic:
		return nil
	}
	return errors.New("invalid platform.network-function type")
}

type NetworkingOperators string

const (
	CNI NetworkingOperators = "cko-cni"
	SM  NetworkingOperators = "cko-sm"
)

func (n NetworkingOperators) IsValid() error {
	switch n {
	case CNI, SM:
		return nil
	}
	return errors.New("invalid networking.network-function type")
}

type OPSubType string

// Define Network Function SubTypes.
const (
	// platform subtype.
	Prometheus OPSubType = "prometheus"
	Argo       OPSubType = "argo"
	CChecker   OPSubType = "connectivity-checker"
	// networking subtype.
	Calico  OPSubType = "calico"
	Cilium  OPSubType = "cilium"
	Aci     OPSubType = "aci-cni"
	SMM     OPSubType = "smm"
	MetalLB OPSubType = "metallb"
)

// InstallerStatus defines the observed state of Installer.
type InstallerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ObservedGeneration int64              `json:"observedGeneration,omitempty"`
	InstallerPhase     Phase              `json:"operatorStatus,omitempty"`
	ComponentStatus    map[string]NFState `json:"componentStatus,omitempty"`
	// store count of nodes, ns, pods & services
	ClusterOverview map[string]int `json:"clusterOverview,omitempty"`
}

type Phase string

const (
	Pending   Phase = "Pending"
	Running   Phase = "Running"
	Succeeded Phase = "Succeeded"
	Failed    Phase = "Failed"
	Unknown   Phase = "Unknown"
)

type NFState struct {
	Type    OPSubType `json:"type,omitempty"`
	Status  string    `json:"status,omitempty"`
	NFState string    `json:"state,omitempty"`
}

type Installer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstallerSpec   `json:"spec,omitempty"`
	Status InstallerStatus `json:"status,omitempty"`
}

// InstallerList contains a list of Installer.
type InstallerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Installer `json:"items"`
}
