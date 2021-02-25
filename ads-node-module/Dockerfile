FROM golang:1.16 AS build
WORKDIR /go/src/
COPY . /go/src/ads-node-module
ENV CGO_ENABLED=0
ENV GOPATH=/go
ARG VERSION=dev

WORKDIR /go/src/ads-node-module
RUN make test && VERSION=${VERSION} BIN_DIR=/install make

FROM scratch
COPY --from=build /install/ads-node-module /ads-node-module
ENTRYPOINT ["/ads-node-module"]
