ARG ARCH="amd64"
ARG OS="linux"
FROM quay.io/prometheus/busybox-${OS}-${ARCH}:latest

COPY ./bin/remote/prometheus-podman3-exporter /bin/podman3_exporter

EXPOSE 9882
USER nobody
ENTRYPOINT [ "/bin/podman3_exporter" ]
