package pdcs

import (
	"strings"

	"github.com/containers/podman/v3/cmd/podman/registry"
	podnetwork "github.com/containers/podman/v3/libpod/network"
	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/containers/podman/v3/pkg/network"
)

// Network implements network's basic information.
type Network struct {
	Name             string
	ID               string
	Driver           string
	NetworkInterface string
	Labels           string
}

type listPrintReports struct {
	*entities.NetworkListReport
}

// Networks returns list of networks (Network).
func Networks() ([]Network, error) {
	networks := make([]Network, 0)

	reports, err := registry.ContainerEngine().NetworkList(registry.Context(), entities.NetworkListOptions{})
	if err != nil {
		return networks, err
	}

	for _, rep := range reports {
		// only two possible values in v3
		driver := podnetwork.DefaultNetworkDriver
		plugins := network.GetCNIPlugins(rep.NetworkConfigList)
		if strings.Contains(plugins, podnetwork.MacVLANNetworkDriver) {
			driver = podnetwork.MacVLANNetworkDriver
		}

		networks = append(networks, Network{
			Name:             rep.Name,
			ID:               getID(network.GetNetworkID(rep.Name)),
			Driver:           driver,
			NetworkInterface: "", // not supported
			Labels:           listPrintReports{rep}.labels(),
		})
	}

	return networks, nil
}

func (n listPrintReports) labels() string {
	list := make([]string, 0, len(n.Labels))
	for k, v := range n.Labels {
		list = append(list, k+"="+v)
	}

	return strings.Join(list, ",")
}
