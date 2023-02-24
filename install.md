## Installation Guide

- [**Building From Source**](#building-from-source)

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
