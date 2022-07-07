FROM golang:1.15-alpine3.12 as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o ProyectoGo

CMD ["/ProyectoGo"]