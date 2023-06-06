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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InstallerSpec defines the desired state of Installer.
type InstallerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// OperatorConfigs describes Platform and Network components to be provisioned and managed
	OperatorConfigs OperatorConfigs `json:"operatorConfigs,omitempty"`
}

type OperatorConfigs struct {
	// Platform will contain a map describing auxiliary applications/features
	Platform map[PlatformOperator]OperatorInfo `json:"platform,omitempty"`
	// Network will contain a map describing Networking operators
	Network map[NetworkOperator]OperatorInfo `json:"network,omitempty"`
}

// OperatorInfo describes OPerators that can be managed by the Installer
type OperatorInfo struct {
	OPSubType OPSubType `json:"type,omitempty"`
	// Manifests will contain base64 encoded specs of the operator to be managed by netop-manager.
	Manifests       map[string]string   `json:"manifests,omitempty"`
	ArrayOptions    map[string][]string `json:"arrayOptions,omitempty"`
	WorkloadVersion string              `json:"version,omitempty"`
	// ManagedComponent field is ignored for platform components and must be set to false to adopt networking components without managing them.
	ManagedComponent bool `json:"managedComponent,omitempty"`
}

// Installer Components.
const Platform = "Platform"
const Network = "Network"

type PlatformOperator string

const (
	Observed   PlatformOperator = "cko-observed"
	Diagnostic PlatformOperator = "cko-diagnostic"
	Git        PlatformOperator = "cko-gitops"
)

type NetworkOperator string

const (
	CNI NetworkOperator = "cko-cni"
	SM  NetworkOperator = "cko-sm"
)

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
	Conditions          []metav1.Condition          `json:"conditions,omitempty"`
	ObservedGeneration  int64                       `json:"observedGeneration,omitempty"`
	InstallerPhase      Phase                       `json:"phase,omitempty"`
	ComponentConditions map[string]metav1.Condition `json:"componentStatus,omitempty"`
}

type Phase string

const (
	Pending     Phase = "Pending"
	Running     Phase = "Running"
	Failed      Phase = "Failed"
	Unknown     Phase = "Unknown"
	Terminating Phase = "Terminating"
)

// Conditions.
const ConditionProvisioned = "Provisioned"
const ConditionProvisionComplete = "Complete"
const ConditionProvisionError = "Error"

// Reasons.
const InstallerReasonProvisionedChildCR = "InstallerProvisionedCR"
const InstallerReasonFailedProvisioningChildCR = "InstallerFailedProvisioningCR"
const InstallerReasonChildStateRunning = "ChildStateRunning"
const InstallerReasonChildStateError = "ChildStateError"

// Annotations and Labels.
const InstallerBase = "netop-manager.io/"
const InstallerComponent = "netop-manager.io/component"
const InstallerComponentType = "netop-manager.io/type"

type NFState struct {
	Type        OPSubType `json:"type,omitempty"`
	Mode        string    `json:"mode,omitempty"`
	LastUpdated string    `json:"lastUpdated,omitempty"`
}

// Installer is the Schema for the installers API
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
