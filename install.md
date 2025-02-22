## Installation Guide

- [**Building From Source**](#building-from-source)
- [**Building with the help of podman**](#building-with-the-help-of-podman)
- [**Container Image**](#container-image)

## Building From Source

prometheus-podman3-exporter is using go v1.17 or above.

1. Clone the repo
2. Install dependencies
    * Fedora

        ```shell
        $ sudo dnf install -y btrfs-progs-devel device-mapper-devel gpgme-devel libassuan-devel
        ```

    * Debian

        ```shell
        $ sudo apt-get -y install libgpgme-dev libbtrfs-dev libdevmapper-dev libassuan-dev pkg-config
        ```

2. Build and run the executable

    ```shell
    $ make binary
    $ ./bin/prometheus-podman3-exporter
    ```

## Building with the help of podman

Keep your environment clean and tidy.

```shell
APP=prometheus-podman3-exporter
podman run --rm -i -v $PWD:/usr/src/$APP -w /usr/src/$APP docker.io/library/golang:1.17-bullseye bash <<EOF
apt-get update
apt-get install -y libgpgme-dev libbtrfs-dev libdevmapper-dev libassuan-dev pkg-config make
make binary
EOF
```
Voilà! Find the binary at ``./bin/prometheus-podman3-exporter``.

## Container Image

It seems safe to use a lower version of client library to connect to podman remotely (in tunnel mode), so an image built with lower version is provided.

* Using unix socket (rootless):

Remember to **enable linger in loginctl**, otherwise your socket will disappear after you log out!

 ```shell
# loginctl enable-linger <username>
 
systemctl start --user podman.socket
podman run -e CONTAINER_HOST=unix:///run/podman/podman.sock -v $XDG_RUNTIME_DIR/podman/podman.sock:/run/podman/podman.sock --userns=keep-id --security-opt label=disable -p 9882:9882 ghcr.io/yylyyl/prometheus-podman3-exporter:latest
 ```

* Using unix socket (root):

 ```shell
systemctl start podman.socket
podman run -e CONTAINER_HOST=unix:///run/podman/podman.sock -v /run/podman/podman.sock:/run/podman/podman.sock --security-opt label=disable -p 9882:9882 ghcr.io/yylyyl/prometheus-podman3-exporter:latest
 ```

* Using TCP:

 ```shell
podman system service --time=0 tcp://<ip>:<port>
podman run -e CONTAINER_HOST=tcp://<ip>:<port> --network=host -p 9882:9882 ghcr.io/yylyyl/prometheus-podman3-exporter:latest
 ```

