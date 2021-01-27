FROM golang as builder
WORKDIR /go/src/github.com/dkeohane/yagsy
# RUN go get github.com/derekparker/delve/cmd/dlv

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o yagsy .

######## Start a new stage from scratch #######
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/github.com/dkeohane/yagsy .

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

EXPOSE 8080 2345