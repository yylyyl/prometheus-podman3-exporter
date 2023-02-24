module github.com/yylyyl/prometheus-podman3-exporter

go 1.17

require (
	github.com/containers/image/v5 v5.11.1
	github.com/containers/podman/v3 v3.1.2
	github.com/cri-o/ocicni v0.2.1-0.20210301205850-541cf7c703cf
	github.com/go-kit/log v0.2.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.14.0
	github.com/prometheus/common v0.39.0
	github.com/prometheus/exporter-toolkit v0.8.2
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.6.1
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/Microsoft/go-winio v0.4.17-0.20210211115548-6eac466e5fa3 // indirect
	github.com/Microsoft/hcsshim v0.8.16 // indirect
	github.com/VividCortex/ewma v1.1.1 // indirect
	github.com/acarl005/stripansi v0.0.0-20180116102854-5a71ef0e047d // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/checkpoint-restore/checkpointctl v0.0.0-20210301084134-a2024f5584e7 // indirect
	github.com/checkpoint-restore/go-criu v0.0.0-20190109184317-bdb7599cd87b // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/containerd/cgroups v0.0.0-20210114181951-8a68de567b68 // indirect
	github.com/containerd/containerd v1.5.0-beta.4 // indirect
	github.com/containerd/continuity v0.0.0-20210208174643-50096c924a4e // indirect
	github.com/containernetworking/cni v0.8.1 // indirect
	github.com/containernetworking/plugins v0.9.1 // indirect
	github.com/containers/buildah v1.20.1 // indirect
	github.com/containers/common v0.36.0 // indirect
	github.com/containers/conmon v2.0.20+incompatible // indirect
	github.com/containers/libtrust v0.0.0-20190913040956-14b96171aa3b // indirect
	github.com/containers/ocicrypt v1.1.0 // indirect
	github.com/containers/psgo v1.5.2 // indirect
	github.com/containers/storage v1.29.0 // indirect
	github.com/coreos/go-iptables v0.5.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/coreos/go-systemd/v22 v22.4.0 // indirect
	github.com/cyphar/filepath-securejoin v0.2.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.3-0.20210216175712-646072ed6524+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-plugins-helpers v0.0.0-20200102110956-c9a8a2d92ccc // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/docker/libnetwork v0.8.0-dev.2.0.20190625141545-5a177b73e316 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/fsouza/go-dockerclient v1.7.2 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v0.2.0 // indirect
	github.com/godbus/dbus/v5 v5.0.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-intervals v0.0.2 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/shlex v0.0.0-20181106134648-c34317bd91bf // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/ishidawataru/sctp v0.0.0-20191218070446-00ab2ac2db07 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/ansiterm v0.0.0-20180109212912-720a0952cc2a // indirect
	github.com/klauspost/compress v1.11.13 // indirect
	github.com/klauspost/pgzip v1.2.5 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/lunixbochs/vtclean v0.0.0-20180621232353-2d01aacdc34a // indirect
	github.com/manifoldco/promptui v0.8.0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mattn/go-shellwords v1.0.11 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/moby/sys/mount v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.4.1 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/mtrmac/gpgme v0.1.2 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2-0.20190823105129-775207bd45b6 // indirect
	github.com/opencontainers/runc v1.0.0-rc93 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20200929063507-e6143ca7d51d // indirect
	github.com/opencontainers/runtime-tools v0.9.0 // indirect
	github.com/opencontainers/selinux v1.8.0 // indirect
	github.com/openshift/imagebuilder v1.2.0 // indirect
	github.com/ostreedev/ostree-go v0.0.0-20190702140239-759a8c1ac913 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/pquerna/ffjson v0.0.0-20190813045741-dac163c6c0a9 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rootless-containers/rootlesskit v0.14.0-beta.0 // indirect
	github.com/safchain/ethtool v0.0.0-20190326074333-42ed695e3de8 // indirect
	github.com/seccomp/libseccomp-golang v0.9.2-0.20200616122406-847368b35ebf // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stefanberger/go-pkcs11uri v0.0.0-20201008174630-78d3cae3a980 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/tchap/go-patricia v2.3.0+incompatible // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	github.com/vbatts/tar-split v0.11.1 // indirect
	github.com/vbauerster/mpb/v6 v6.0.3 // indirect
	github.com/vishvananda/netlink v1.1.1-0.20201029203352-d40f9887b852 // indirect
	github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae // indirect
	github.com/willf/bitset v1.1.11 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190809123943-df4f5c81cb3b // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.mozilla.org/pkcs7 v0.0.0-20200128120323-432b2356ecb1 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/oauth2 v0.3.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/term v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220616135557-88e70c0c3a90 // indirect
	google.golang.org/grpc v1.47.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.20.1 // indirect
	k8s.io/apimachinery v0.20.5 // indirect
	k8s.io/klog/v2 v2.4.0 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.0.2 // indirect
)
