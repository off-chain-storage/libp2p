FROM golang:1.21.2-alpine as builder
RUN apk update

WORKDIR /usr/src/app
COPY . .

ENV GO111MODULE=on

# RUN go get -u github.com/off-chain-storage/go-off-chain-storage@v1.2.1
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o bin/main main/main.go

### Executable Image
FROM alpine

COPY --from=builder /usr/src/app/bin/main ./main

ENTRYPOINT ["./main"]