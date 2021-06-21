FROM golang:1.15.8 as builder
ADD . /build
WORKDIR /build
RUN go vet ./...
RUN go test ./...
RUN go build -o build/crumblecog-plugin

FROM alpine as putter
COPY --from=builder /build/build/crumblecog-plugin .
ENTRYPOINT [ "mv", "crumblecog-plugin", "/custom-tools/" ]