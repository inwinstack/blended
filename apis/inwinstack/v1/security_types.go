/*
Copyright © 2018 inwinSTACK.inc

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
	Type                            string
	Description                     string              `json:"description"`
	SourceZones                     []string            `json:"sourceZones"`
	SourceAddresses                 []string            `json:"sourceAddresses"`
	SourceUsers                     []string            `json:"sourceUsers"`
	HipProfiles                     []string            `json:"hipProfiles"`
	DestinationZones                []string            `json:"destinationZones"`
	DestinationAddresses            []string            `json:"destinationAddresses"`
	Applications                    []string            `json:"applications"`
	Services                        []string            `json:"services"`
	Categories                      []string            `json:"categories"`
	Action                          string              `json:"action"`
	IcmpUnreachable                 bool                `json:"icmpUnreachable"`
	DisableServerResponseInspection bool                `json:"disableServerResponseInspection"`
	NegateDestination               bool                `json:"negateDestination"`
	NegateSource                    bool                `json:"negateSource"`
	NegateTarget                    bool                `json:"negateTarget"`
	LogSetting                      string              `json:"logSetting"`
	LogStart                        bool                `json:"logStart"`
	LogEnd                          bool                `json:"logEnd"`
	Disabled                        bool                `json:"disabled"`
	Schedule                        string              `json:"schedule"`
	Group                           string              `json:"group"`
	Targets                         map[string][]string `json:"targets"`
	Virus                           string              `json:"virus"`
	Spyware                         string              `json:"spyware"`
	Vulnerability                   string              `json:"vulnerability"`
	URLFiltering                    string              `json:"urlFiltering"`
	FileBlocking                    string              `json:"fileBlocking"`
	WildFireAnalysis                string              `json:"wildFireAnalysis"`
	DataFiltering                   string              `json:"dataFiltering"`
	Tags                            []string            `json:"tags"`
}

type SecurityPhase string

// These are the valid phases of a security.
const (
	SecurityPending SecurityPhase = "Pending"
	SecurityActive  SecurityPhase = "Active"
	SecurityFailed  SecurityPhase = "Failed"
)

// SecurityStatus represents the current state of a security resource.
type SecurityStatus struct {
	Phase          SecurityPhase `json:"phase"`
	Reason         string        `json:"reason,omitempty"`
	LastUpdateTime metav1.Time   `json:"lastUpdateTime"`
}
