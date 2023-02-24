// +build linux

package buildah

import (
	"github.com/opencontainers/runtime-tools/generate"
	selinux "github.com/opencontainers/selinux/go-selinux"
)

func selinuxGetEnabled() bool {
	return selinux.GetEnabled()
}

func setupSelinux(g *generate.Generator, processLabel, mountLabel string) {
	if processLabel != "" && selinux.GetEnabled() {
		g.SetProcessSelinuxLabel(processLabel)
		g.SetLinuxMountLabel(mountLabel)
	}
}
