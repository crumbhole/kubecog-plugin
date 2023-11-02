FROM golang:1.21.3 as builder
COPY --from=hairyhenderson/gomplate:v3.11.5-alpine /bin/gomplate /usr/local/bin/gomplate
ADD . /build
WORKDIR /build
RUN go vet ./...
RUN go test ./...
RUN go build -buildvcs=false -o build/kubecog-plugin

FROM ghcr.io/crumbhole/argocd-lovely-plugin-cmp-vault:0.18.0 as putter
COPY --from=builder /build/build/kubecog-plugin /usr/local/bin
COPY --from=builder /usr/local/bin/gomplate /usr/local/bin
