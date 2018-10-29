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
	Type                           string              `json:"type"`
	Description                    string              `json:"description"`
	SourceZones                    []string            `json:"sourceZones"`
	SourceAddresses                []string            `json:"sourceAddresses"`
	DestinationAddresses           []string            `json:"destinationAddresses"`
	DestinationZone                string              `json:"destinationZone"`
	ToInterface                    string              `json:"toInterface"`
	Service                        string              `json:"service"`
	SatType                        string              `json:"satType"`
	SatAddressType                 string              `json:"satAddressType"`
	SatTranslatedAddresses         []string            `json:"satTranslatedAddresses"`
	SatInterface                   string              `json:"satInterface"`
	SatIPAddress                   string              `json:"satIPAddress"`
	SatFallbackType                string              `json:"satFallbackType"`
	SatFallbackTranslatedAddresses []string            `json:"satFallbackTranslatedAddresses"`
	SatFallbackInterface           string              `json:"satFallbackInterface"`
	SatFallbackIPType              string              `json:"satFallbackIPType"`
	SatFallbackIPAddress           string              `json:"satFallbackIPAddress"`
	SatStaticTranslatedAddress     string              `json:"satStaticTranslatedAddress"`
	SatStaticBiDirectional         bool                `json:"satStaticBiDirectional"`
	DatType                        string              `json:"datType"`
	DatAddress                     string              `json:"datAddress"`
	DatPort                        int32               `json:"datPort"`
	DatDynamicDistribution         string              `json:"datDynamicDistribution"`
	Disabled                       bool                `json:"disabled"`
	Targets                        map[string][]string `json:"targets"`
	NegateTarget                   bool                `json:"negateTarget"`
	Tags                           []string            `json:"tags"`
}

type NATPhase string

// These are the valid phases of a NAT.
const (
	NATPending NATPhase = "Pending"
	NATActive  NATPhase = "Active"
	NATFailed  NATPhase = "Failed"
)

// NATStatus represents the current state of a NAT resource.
type NATStatus struct {
	Phase          NATPhase    `json:"phase"`
	Reason         string      `json:"reason,omitempty"`
	LastUpdateTime metav1.Time `json:"lastUpdateTime"`
}
