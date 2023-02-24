package filters

import (
	"strconv"
	"strings"
	"time"

	"github.com/containers/podman/v3/libpod"
	"github.com/containers/podman/v3/libpod/define"
	"github.com/containers/podman/v3/pkg/network"
	"github.com/containers/podman/v3/pkg/util"
	"github.com/pkg/errors"
)

// GenerateContainerFilterFuncs return ContainerFilter functions based of filter.
func GenerateContainerFilterFuncs(filter string, filterValues []string, r *libpod.Runtime) (func(container *libpod.Container) bool, error) {
	switch filter {
	case "id":
		// we only have to match one ID
		return func(c *libpod.Container) bool {
			return util.StringMatchRegexSlice(c.ID(), filterValues)
		}, nil
	case "label":
		// we have to match that all given labels exits on that container
		return func(c *libpod.Container) bool {
			return util.MatchLabelFilters(filterValues, c.Labels())
		}, nil
	case "name":
		// we only have to match one name
		return func(c *libpod.Container) bool {
			return util.StringMatchRegexSlice(c.Name(), filterValues)
		}, nil
	case "exited":
		var exitCodes []int32
		for _, exitCode := range filterValues {
			ec, err := strconv.ParseInt(exitCode, 10, 32)
			if err != nil {
				return nil, errors.Wrapf(err, "exited code out of range %q", ec)
			}
			exitCodes = append(exitCodes, int32(ec))
		}
		return func(c *libpod.Container) bool {
			ec, exited, err := c.ExitCode()
			if err == nil && exited {
				for _, exitCode := range exitCodes {
					if ec == exitCode {
						return true
					}
				}
			}
			return false
		}, nil
	case "status":
		for _, filterValue := range filterValues {
			if !util.StringInSlice(filterValue, []string{"created", "running", "paused", "stopped", "exited", "unknown"}) {
				return nil, errors.Errorf("%s is not a valid status", filterValue)
			}
		}
		return func(c *libpod.Container) bool {
			status, err := c.State()
			if err != nil {
				return false
			}
			state := status.String()
			if status == define.ContainerStateConfigured {
				state = "created"
			} else if status == define.ContainerStateStopped {
				state = "exited"
			}
			for _, filterValue := range filterValues {
				if filterValue == "stopped" {
					filterValue = "exited"
				}
				if state == filterValue {
					return true
				}
			}
			return false
		}, nil
	case "ancestor":
		// This needs to refine to match docker
		// - ancestor=(<image-name>[:tag]|<image-id>| ⟨image@digest⟩) - containers created from an image or a descendant.
		return func(c *libpod.Container) bool {
			for _, filterValue := range filterValues {
				containerConfig := c.Config()
				if strings.Contains(containerConfig.RootfsImageID, filterValue) || strings.Contains(containerConfig.RootfsImageName, filterValue) {
					return true
				}
			}
			return false
		}, nil
	case "before":
		var createTime time.Time
		for _, filterValue := range filterValues {
			ctr, err := r.LookupContainer(filterValue)
			if err != nil {
				return nil, err
			}
			containerConfig := ctr.Config()
			if createTime.IsZero() || createTime.After(containerConfig.CreatedTime) {
				createTime = containerConfig.CreatedTime
			}
		}
		return func(c *libpod.Container) bool {
			cc := c.Config()
			return createTime.After(cc.CreatedTime)
		}, nil
	case "since":
		var createTime time.Time
		for _, filterValue := range filterValues {
			ctr, err := r.LookupContainer(filterValue)
			if err != nil {
				return nil, err
			}
			containerConfig := ctr.Config()
			if createTime.IsZero() || createTime.After(containerConfig.CreatedTime) {
				createTime = containerConfig.CreatedTime
			}
		}
		return func(c *libpod.Container) bool {
			cc := c.Config()
			return createTime.Before(cc.CreatedTime)
		}, nil
	case "volume":
		//- volume=(<volume-name>|<mount-point-destination>)
		return func(c *libpod.Container) bool {
			containerConfig := c.Config()
			var dest string
			for _, filterValue := range filterValues {
				arr := strings.SplitN(filterValue, ":", 2)
				source := arr[0]
				if len(arr) == 2 {
					dest = arr[1]
				}
				for _, mount := range containerConfig.Spec.Mounts {
					if dest != "" && (mount.Source == source && mount.Destination == dest) {
						return true
					}
					if dest == "" && mount.Source == source {
						return true
					}
				}
				for _, vname := range containerConfig.NamedVolumes {
					if dest != "" && (vname.Name == source && vname.Dest == dest) {
						return true
					}
					if dest == "" && vname.Name == source {
						return true
					}
				}
			}
			return false
		}, nil
	case "health":
		return func(c *libpod.Container) bool {
			hcStatus, err := c.HealthCheckStatus()
			if err != nil {
				return false
			}
			for _, filterValue := range filterValues {
				if hcStatus == filterValue {
					return true
				}
			}
			return false
		}, nil
	case "until":
		return prepareUntilFilterFunc(filterValues)
	case "pod":
		var pods []*libpod.Pod
		for _, podNameOrID := range filterValues {
			p, err := r.LookupPod(podNameOrID)
			if err != nil {
				if errors.Cause(err) == define.ErrNoSuchPod {
					continue
				}
				return nil, err
			}
			pods = append(pods, p)
		}
		return func(c *libpod.Container) bool {
			// if no pods match, quick out
			if len(pods) < 1 {
				return false
			}
			// if the container has no pod id, quick out
			if len(c.PodID()) < 1 {
				return false
			}
			for _, p := range pods {
				// we already looked up by name or id, so id match
				// here is ok
				if p.ID() == c.PodID() {
					return true
				}
			}
			return false
		}, nil
	case "network":
		return func(c *libpod.Container) bool {
			networks, _, err := c.Networks()
			// if err or no networks, quick out
			if err != nil || len(networks) == 0 {
				return false
			}
			for _, net := range networks {
				netID := network.GetNetworkID(net)
				for _, val := range filterValues {
					// match by network name or id
					if val == net || val == netID {
						return true
					}
				}
			}
			return false
		}, nil
	}
	return nil, errors.Errorf("%s is an invalid filter", filter)
}

// GeneratePruneContainerFilterFuncs return ContainerFilter functions based of filter for prune operation
func GeneratePruneContainerFilterFuncs(filter string, filterValues []string, r *libpod.Runtime) (func(container *libpod.Container) bool, error) {
	switch filter {
	case "label":
		return func(c *libpod.Container) bool {
			return util.MatchLabelFilters(filterValues, c.Labels())
		}, nil
	case "until":
		return prepareUntilFilterFunc(filterValues)
	}
	return nil, errors.Errorf("%s is an invalid filter", filter)
}

func prepareUntilFilterFunc(filterValues []string) (func(container *libpod.Container) bool, error) {
	until, err := util.ComputeUntilTimestamp(filterValues)
	if err != nil {
		return nil, err
	}
	return func(c *libpod.Container) bool {
		if !until.IsZero() && c.CreatedTime().Before(until) {
			return true
		}
		return false
	}, nil
}
