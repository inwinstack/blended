/*
Copyright © 2018 inwinSTACK Inc

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

package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceList is a list of service.
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Service `json:"items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Service represents a Kubernetes Service Custom Resource.
// The Service will be used as PA service object.
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ServiceSpec   `json:"spec"`
	Status ServiceStatus `json:"status,omitempty"`
}

// ServiceSpec is the spec for a Service resource.
type ServiceSpec struct {
	Description     string   `json:"description,omitempty"`
	Protocol        string   `json:"protocol,omitempty"`
	SourcePort      string   `json:"sourcePort,omitempty"`
	DestinationPort string   `json:"destinationPort,omitempty"`
	Tags            []string `json:"tags,omitempty"`
}

type ServicePhase string

// These are the valid phases of a Service.
const (
	ServiceNone        ServicePhase = ""
	ServicePending     ServicePhase = "Pending"
	ServiceActive      ServicePhase = "Active"
	ServiceFailed      ServicePhase = "Failed"
	ServiceTerminating ServicePhase = "Terminating"
)

// ServiceStatus represents the current state of a Service resource.
type ServiceStatus struct {
	Phase          ServicePhase `json:"phase"`
	Reason         string       `json:"reason,omitempty"`
	LastUpdateTime metav1.Time  `json:"lastUpdateTime"`
}
