FROM golang:1.11.4

WORKDIR /go/src/peercast-yayp/

RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY . ./
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o yayp

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /go/src/peercast-yayp/yayp .
CMD ["./yayp"]
