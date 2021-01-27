FROM golang:alpine AS builder
WORKDIR /go/src/pod-controller
COPY . .
RUN cd cmd \
    && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o  pod-controller .

FROM scratch
COPY --from=builder /go/src/pod-controller/cmd/pod-controller /bin/pod-controller

ENTRYPOINT ["/bin/pod-controller"]
