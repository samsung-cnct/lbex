package nginx

import (
	"k8s.io/client-go/pkg/api/v1"
)

var (
	SupportedAlgorithms = []string{
		"roundrobin", // *set as default below* direct traffic sequentially to the servers.
		"leastconn",  // selects the server with the smaller number of current active connections.
		"leasttime",  // selects the server with the lowest average latency and the least number of active connections.
	}
	DefaultAlgorithm = SupportedAlgorithms[0]

	SupportedMethods = []string{
		"connect", // *set as default value below*
		"first_byte",
		"last_byte",
		"connect inflight",
		"first_byte inflight",
		"last_byte inflight",
	}
	DefaultMethod = SupportedMethods[0]

	// UpstreamTypes - service upstream pool target types
	UpstreamTypes = []string{
		"node", // *set as default value below*
		"pod",
		"cluster-ip",
	}

	// DefaultUpstreamType - default service upstream pool target type
	DefaultUpstreamType = UpstreamTypes[0]
)

// Target is a service network topology target
type Target struct {
	// ServicePort - the port that we listen on for the service's external clients
	ServicePort int
	// NodeIP - the IP address of a host/worker node
	NodeIP string
	// NodeName - the name of a host/worker node
	NodeName string
	// NodePort - the port that the host/worker node listens on for fowarding to the pod/ip:port
	NodePort int
	// PortName - the name of the port if present, or 'unnamed' otherwise
	PortName string
	// PodIP - the pods ip address
	PodIP string
	// PodPort - the port the that the pod listens on
	PodPort int
	// Protocol - TCP or UDP
	Protocol string
}

// ServiceSpec models basic Service details and the Endpoints of the services
type ServiceSpec struct {
	Service      *v1.Service
	Key          string
	Algorithm    string
	ClusterIP    string
	ConfigName   string
	UpstreamType string
	Topology     []Target
}

// ValidateAlgorithm - returns the input 'a' algorithm value iff it is a valid
// value from SupportedAlgorithms, otherwise returns default algorithm value
func ValidateAlgorithm(a string) string {
	found := false
	for _, current := range SupportedAlgorithms {
		if a == current {
			found = true
			break
		}
	}
	if !found {
		return DefaultAlgorithm
	}
	return a
}

// ValidateMethod - returns the input 'm' method value iff it is a valid value
// from SupportedMethods, otherwise returns default method value
func ValidateMethod(m string) string {
	found := false
	for _, current := range SupportedMethods {
		if m == current {
			found = true
			break
		}
	}
	if !found {
		return DefaultMethod
	}
	return m
}

// ValidateUpstreamType - returns the input 'ups' upstream type iff it is a
// valid value from UpstreamTypes, otherwise returns default upstream type
func ValidateUpstreamType(ups string) string {
	found := false
	for _, current := range UpstreamTypes {
		if ups == current {
			found = true
			break
		}
	}
	if !found {
		return DefaultUpstreamType
	}
	return ups
}
