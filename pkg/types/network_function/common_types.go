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
	Options   map[string]string `json:"options,omitempty"`
	Version   string            `json:"version"`
	Unmanaged bool              `json:"unmanaged,omitempty"`
}

type OpsState string

const (
	OpsPending      OpsState = "Pending"
	OpsAddFinalizer OpsState = "AddingFinalizer"
	OpsRunnnig      OpsState = "Running"
	OpsSucceeded    OpsState = "Succeeded"
	OpsFailed       OpsState = "Failed"
)
