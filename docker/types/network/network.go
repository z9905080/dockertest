// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package network // import "github.com/z9905080/dockertest/v3/docker/types/network"

// Address represents an IP address
type Address struct {
	Addr      string
	PrefixLen int
}

// IPAM represents IP Address Management
type IPAM struct {
	Driver  string
	Options map[string]string //Per network IPAM driver options
	Config  []IPAMConfig
}

// IPAMConfig represents IPAM configurations
type IPAMConfig struct {
	Subnet     string            `json:",omitempty"`
	IPRange    string            `json:",omitempty"`
	Gateway    string            `json:",omitempty"`
	AuxAddress map[string]string `json:"AuxiliaryAddresses,omitempty"`
}

// EndpointIPAMConfig represents IPAM configurations for the endpoint
type EndpointIPAMConfig struct {
	IPv4Address  string   `json:",omitempty"`
	IPv6Address  string   `json:",omitempty"`
	LinkLocalIPs []string `json:",omitempty"`
}

// Copy makes a copy of the endpoint ipam config
func (cfg *EndpointIPAMConfig) Copy() *EndpointIPAMConfig {
	cfgCopy := *cfg
	cfgCopy.LinkLocalIPs = make([]string, 0, len(cfg.LinkLocalIPs))
	cfgCopy.LinkLocalIPs = append(cfgCopy.LinkLocalIPs, cfg.LinkLocalIPs...)
	return &cfgCopy
}

// PeerInfo represents one peer of an overlay network
type PeerInfo struct {
	Name string
	IP   string
}

// EndpointSettings stores the network endpoint details
type EndpointSettings struct {
	// Configurations
	IPAMConfig *EndpointIPAMConfig
	Links      []string
	Aliases    []string
	// Operational data
	NetworkID           string
	EndpointID          string
	Gateway             string
	IPAddress           string
	IPPrefixLen         int
	IPv6Gateway         string
	GlobalIPv6Address   string
	GlobalIPv6PrefixLen int
	MacAddress          string
	DriverOpts          map[string]string
}

// Task carries the information about one backend task
type Task struct {
	Name       string
	EndpointID string
	EndpointIP string
	Info       map[string]string
}

// ServiceInfo represents service parameters with the list of service's tasks
type ServiceInfo struct {
	VIP          string
	Ports        []string
	LocalLBIndex int
	Tasks        []Task
}

// Copy makes a deep copy of `EndpointSettings`
func (es *EndpointSettings) Copy() *EndpointSettings {
	epCopy := *es
	if es.IPAMConfig != nil {
		epCopy.IPAMConfig = es.IPAMConfig.Copy()
	}

	if es.Links != nil {
		links := make([]string, 0, len(es.Links))
		epCopy.Links = append(links, es.Links...)
	}

	if es.Aliases != nil {
		aliases := make([]string, 0, len(es.Aliases))
		epCopy.Aliases = append(aliases, es.Aliases...)
	}
	return &epCopy
}

// NetworkingConfig represents the container's networking configuration for each of its interfaces
// Carries the networking configs specified in the `docker run` and `docker network connect` commands
type NetworkingConfig struct {
	EndpointsConfig map[string]*EndpointSettings // Endpoint configs for each connecting network
}

// ConfigReference specifies the source which provides a network's configuration
type ConfigReference struct {
	Network string
}
