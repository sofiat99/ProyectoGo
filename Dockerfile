FROM golang:1.15-alpine3.12 as builder
WORKDIR $GOPATH/src/github.com/sofiat99/ProyectoGo
COPY ./ .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v
RUN cp   ./

FROM alpine:3.12
COPY --from=builder / /
CMD ["/ProyectoGo"]