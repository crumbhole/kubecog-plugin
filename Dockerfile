FROM golang:1.20.0 as builder
COPY --from=hairyhenderson/gomplate:v3.11.3-alpine /bin/gomplate /usr/local/bin/gomplate
ADD . /build
WORKDIR /build
RUN go vet ./...
RUN go test ./...
RUN go build -buildvcs=false -o build/kubecog-plugin

FROM alpine:3.17.2 as putter
COPY --from=builder /build/build/kubecog-plugin .
COPY --from=builder /usr/local/bin/gomplate .
USER 999
ENTRYPOINT [ "cp", "kubecog-plugin", "gomplate", "/custom-tools/" ]
