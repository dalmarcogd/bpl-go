FROM golang:1.15.2-buster as builder

WORKDIR /go/src/github.com/dalmarcogd/bpl-go/
COPY ./ /go/src/github.com/dalmarcogd/bpl-go/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GODEBUG=madvdontneed=1 go build -a -tags netgo -o application cmd/api/main.go && mv application /application

FROM alpine:3.12.0 as runner
WORKDIR /
COPY --from=builder /application .
EXPOSE 8080
ENTRYPOINT ["./application"]