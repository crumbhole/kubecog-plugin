FROM golang:1.20.2 as builder
COPY --from=hairyhenderson/gomplate:v3.11.4-alpine /bin/gomplate /usr/local/bin/gomplate
ADD . /build
WORKDIR /build
RUN go vet ./...
RUN go test ./...
RUN go build -buildvcs=false -o build/kubecog-plugin

FROM alpine:3.17.3 as putter
COPY --from=builder /build/build/kubecog-plugin .
COPY --from=builder /usr/local/bin/gomplate .
USER 999
ENTRYPOINT [ "cp", "kubecog-plugin", "gomplate", "/custom-tools/" ]
