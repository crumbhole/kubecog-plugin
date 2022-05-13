FROM golang:1.18.1 as builder
ADD . /build
WORKDIR /build
RUN go vet ./...
RUN go test ./...
RUN go build -buildvcs=false -o build/crumblecog-plugin

FROM alpine as putter
COPY --from=builder /build/build/crumblecog-plugin .
ENTRYPOINT [ "mv", "crumblecog-plugin", "/custom-tools/" ]