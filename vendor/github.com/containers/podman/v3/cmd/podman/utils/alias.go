package utils

import "github.com/spf13/pflag"

// AliasFlags is a function to handle backwards compatibility with old flags
func AliasFlags(f *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "healthcheck-command":
		name = "health-cmd"
	case "healthcheck-interval":
		name = "health-interval"
	case "healthcheck-retries":
		name = "health-retries"
	case "healthcheck-start-period":
		name = "health-start-period"
	case "healthcheck-timeout":
		name = "health-timeout"
	case "net":
		name = "network"
	case "timeout":
		name = "time"
	case "namespace":
		name = "ns"
	case "storage":
		name = "external"
	case "purge":
		name = "rm"
	case "override-arch":
		name = "arch"
	case "override-os":
		name = "os"
	case "override-variant":
		name = "variant"
	}
	return pflag.NormalizedName(name)
}
