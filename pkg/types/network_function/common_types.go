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

type Config struct {
	// First key should be internal child api
	Workload WorkloadInfo `json:"workload,omitempty"`
}

type WorkloadInfo struct {
	Type      string            `json:"type"`
	Manifests map[string]string `json:"manifests,omitempty"`
	Version   string            `json:"version"`
	//+kubebuilder:default=false
	ManagedComponent bool `json:"managedComponent,omitempty"`
}

type State string

const (
	Initializing State = "Initializing"
	Pending      State = "Pending"
	Finalizer    State = "AddingFinalizer"
	Running      State = "Running"
	Failed       State = "Failed"
	Unknown      State = "Unknown"
)

// Conditions.
const ConditionReady = "Updating"
const ConditionProgressing = "Progressing"
const ConditionAvailable = "Available"
const ConditionDegraded = "Degraded"

// Reasons.
const ReasonProvisioningOPs = "ProvisioningOperator"
const ReasonFailedProvisioningOps = "FailedProvisioningOperator"
const ReasonProvisionedOps = "ProvisionedOperator"
const ReasonOpsStateRunning = "OperatorRunning"
const ReasonFailedMonitoringOps = "FailedMonitoringOperator"

// cniNames.
const (
	CNIMultus   = "multus"
	CNICilium   = "cko-cni-cilium"
	CNICalico   = "cko-cni-calico"
	CNIAci      = "cko-cni-aci"
	CNIAWSVPC   = "cko-cni-awsvpc"
	CNIOVN      = "cko-cni-ovn"
	NetopCNI    = "network-operator-cni"
	CNINotFound = "not-found"
)
