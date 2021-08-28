FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

USER root
COPY *.go ./
COPY main ./main
COPY templates ./templates
COPY paths.json ./

RUN apk add build-base
RUN go build -o /urlshort ./main/main.go

EXPOSE 8080

CMD [ "/urlshort" ]