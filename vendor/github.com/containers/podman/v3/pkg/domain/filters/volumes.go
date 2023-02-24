package filters

import (
	"net/url"
	"strings"

	"github.com/containers/podman/v3/libpod"
	"github.com/containers/podman/v3/pkg/util"
	"github.com/pkg/errors"
)

func GenerateVolumeFilters(filters url.Values) ([]libpod.VolumeFilter, error) {
	var vf []libpod.VolumeFilter
	for filter, v := range filters {
		for _, val := range v {
			switch filter {
			case "name":
				nameVal := val
				vf = append(vf, func(v *libpod.Volume) bool {
					return nameVal == v.Name()
				})
			case "driver":
				driverVal := val
				vf = append(vf, func(v *libpod.Volume) bool {
					return v.Driver() == driverVal
				})
			case "scope":
				scopeVal := val
				vf = append(vf, func(v *libpod.Volume) bool {
					return v.Scope() == scopeVal
				})
			case "label":
				filter := val
				vf = append(vf, func(v *libpod.Volume) bool {
					return util.MatchLabelFilters([]string{filter}, v.Labels())
				})
			case "opt":
				filterArray := strings.SplitN(val, "=", 2)
				filterKey := filterArray[0]
				var filterVal string
				if len(filterArray) > 1 {
					filterVal = filterArray[1]
				} else {
					filterVal = ""
				}
				vf = append(vf, func(v *libpod.Volume) bool {
					for labelKey, labelValue := range v.Options() {
						if labelKey == filterKey && (filterVal == "" || labelValue == filterVal) {
							return true
						}
					}
					return false
				})
			case "dangling":
				danglingVal := val
				invert := false
				switch strings.ToLower(danglingVal) {
				case "true", "1":
					// Do nothing
				case "false", "0":
					// Dangling=false requires that we
					// invert the result of IsDangling.
					invert = true
				default:
					return nil, errors.Errorf("%q is not a valid value for the \"dangling\" filter - must be true or false", danglingVal)
				}
				vf = append(vf, func(v *libpod.Volume) bool {
					dangling, err := v.IsDangling()
					if err != nil {
						return false
					}
					if invert {
						return !dangling
					}
					return dangling
				})
			default:
				return nil, errors.Errorf("%q is an invalid volume filter", filter)
			}
		}
	}
	return vf, nil
}

func GeneratePruneVolumeFilters(filters url.Values) ([]libpod.VolumeFilter, error) {
	var vf []libpod.VolumeFilter
	for filter, v := range filters {
		for _, val := range v {
			switch filter {
			case "label":
				filter := val
				vf = append(vf, func(v *libpod.Volume) bool {
					return util.MatchLabelFilters([]string{filter}, v.Labels())
				})
			default:
				return nil, errors.Errorf("%q is an invalid volume filter", filter)
			}
		}
	}
	return vf, nil
}
