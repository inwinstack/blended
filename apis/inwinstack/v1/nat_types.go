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

// NATList is a list of NAT.
type NATList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NAT `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NAT represents a Kubernetes NAT Custom Resource.
// The NAT will be used as PA NAT policy.
type NAT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   NATSpec   `json:"spec"`
	Status NATStatus `json:"status,omitempty"`
}

// These are the valid values for SatType.
const (
	NATDynamicIPAndPort = "dynamic-ip-and-port"
	NATDynamicIP        = "dynamic-ip"
	NATStaticIP         = "static-ip"
)

// These are the valid values for SatAddressType.
const (
	NATInterfaceAddress  = "interface-address"
	NATTranslatedAddress = "translated-address"
)

// These are the valid settings for SatFallbackIPType.
const (
	NATIP         = "ip"
	NATFloatingIP = "floating"
)

// These are the valid ip type of a NAT.
const (
	NATIPv4 = "ipv4"
	NATIPv6 = "ipv6"
)

// These are the valid sat type of a NAT.
const (
	NATSatNone = "none"
)

// These are the valid dat type of a NAT.
const (
	NATDatStatic  = "destination-translation"
	NATDatDynamic = "dynamic-destination-translation"
)

// NATSpec is the spec for a NAT resource.
type NATSpec struct {
	Type                           string              `json:"type,omitempty"`
	Description                    string              `json:"description,omitempty"`
	SourceZones                    []string            `json:"sourceZones,omitempty"`
	SourceAddresses                []string            `json:"sourceAddresses,omitempty"`
	DestinationAddresses           []string            `json:"destinationAddresses,omitempty"`
	DestinationZone                string              `json:"destinationZone,omitempty"`
	ToInterface                    string              `json:"toInterface,omitempty"`
	Service                        string              `json:"service,omitempty"`
	SatType                        string              `json:"satType,omitempty"`
	SatAddressType                 string              `json:"satAddressType,omitempty"`
	SatTranslatedAddresses         []string            `json:"satTranslatedAddresses,omitempty"`
	SatInterface                   string              `json:"satInterface,omitempty"`
	SatIPAddress                   string              `json:"satIPAddress,omitempty"`
	SatFallbackType                string              `json:"satFallbackType,omitempty"`
	SatFallbackTranslatedAddresses []string            `json:"satFallbackTranslatedAddresses,omitempty"`
	SatFallbackInterface           string              `json:"satFallbackInterface,omitempty"`
	SatFallbackIPType              string              `json:"satFallbackIPType,omitempty"`
	SatFallbackIPAddress           string              `json:"satFallbackIPAddress,omitempty"`
	SatStaticTranslatedAddress     string              `json:"satStaticTranslatedAddress,omitempty"`
	SatStaticBiDirectional         bool                `json:"satStaticBiDirectional,omitempty"`
	DatType                        string              `json:"datType,omitempty"`
	DatAddress                     string              `json:"datAddress,omitempty"`
	DatPort                        int32               `json:"datPort,omitempty"`
	DatDynamicDistribution         string              `json:"datDynamicDistribution,omitempty"`
	Disabled                       bool                `json:"disabled,omitempty"`
	Targets                        map[string][]string `json:"targets,omitempty"`
	NegateTarget                   bool                `json:"negateTarget,omitempty"`
	Tags                           []string            `json:"tags,omitempty"`
}

type NATPhase string

// These are the valid phases of a NAT.
const (
	NATNone        NATPhase = ""
	NATPending     NATPhase = "Pending"
	NATActive      NATPhase = "Active"
	NATFailed      NATPhase = "Failed"
	NATTerminating NATPhase = "Terminating"
)

// NATStatus represents the current state of a NAT resource.
type NATStatus struct {
	Phase          NATPhase    `json:"phase"`
	Reason         string      `json:"reason,omitempty"`
	LastUpdateTime metav1.Time `json:"lastUpdateTime"`
}
