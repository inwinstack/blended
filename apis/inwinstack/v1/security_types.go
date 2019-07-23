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

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecurityList is a list of security.
type SecurityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Security `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Security represents a Kubernetes Security Custom Resource.
// The Security will be used as PA security policy.
type Security struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   SecuritySpec   `json:"spec"`
	Status SecurityStatus `json:"status,omitempty"`
}

// These are the valid default of a security.
const (
	SecurityUniversal = "universal"
)

// These are the valid action of a security.
const (
	SecurityAllow = "allow"
	SecurityDeny  = "deny"
)

// SecuritySpec is the spec for a security resource.
type SecuritySpec struct {
	Type                            string              `json:"type,omitempty"`
	Description                     string              `json:"description,omitempty"`
	SourceZones                     []string            `json:"sourceZones,omitempty"`
	SourceAddresses                 []string            `json:"sourceAddresses,omitempty"`
	SourceUsers                     []string            `json:"sourceUsers,omitempty"`
	HipProfiles                     []string            `json:"hipProfiles,omitempty"`
	DestinationZones                []string            `json:"destinationZones,omitempty"`
	DestinationAddresses            []string            `json:"destinationAddresses,omitempty"`
	Applications                    []string            `json:"applications,omitempty"`
	Services                        []string            `json:"services,omitempty"`
	Categories                      []string            `json:"categories,omitempty"`
	Action                          string              `json:"action,omitempty"`
	IcmpUnreachable                 bool                `json:"icmpUnreachable,omitempty"`
	DisableServerResponseInspection bool                `json:"disableServerResponseInspection"`
	NegateDestination               bool                `json:"negateDestination,omitempty"`
	NegateSource                    bool                `json:"negateSource,omitempty"`
	NegateTarget                    bool                `json:"negateTarget,omitempty"`
	LogSetting                      string              `json:"logSetting,omitempty"`
	LogStart                        bool                `json:"logStart,omitempty"`
	LogEnd                          bool                `json:"logEnd,omitempty"`
	Disabled                        bool                `json:"disabled,omitempty"`
	Schedule                        string              `json:"schedule,omitempty"`
	Group                           string              `json:"group,omitempty"`
	Targets                         map[string][]string `json:"targets,omitempty"`
	Virus                           string              `json:"virus,omitempty"`
	Spyware                         string              `json:"spyware,omitempty"`
	Vulnerability                   string              `json:"vulnerability,omitempty"`
	URLFiltering                    string              `json:"urlFiltering,omitempty"`
	FileBlocking                    string              `json:"fileBlocking,omitempty"`
	WildFireAnalysis                string              `json:"wildFireAnalysis,omitempty"`
	DataFiltering                   string              `json:"dataFiltering,omitempty"`
	Tags                            []string            `json:"tags,omitempty"`
}

type SecurityPhase string

// These are the valid phases of a security.
const (
	SecurityNone        SecurityPhase = ""
	SecurityPending     SecurityPhase = "Pending"
	SecurityActive      SecurityPhase = "Active"
	SecurityFailed      SecurityPhase = "Failed"
	SecurityTerminating SecurityPhase = "Terminating"
)

// SecurityStatus represents the current state of a security resource.
type SecurityStatus struct {
	Phase          SecurityPhase `json:"phase"`
	Reason         string        `json:"reason,omitempty"`
	LastUpdateTime metav1.Time   `json:"lastUpdateTime"`
}
