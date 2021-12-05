FROM golang:1.16-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o /godocker-redis

EXPOSE 8001

CMD ["/godocker-redis"]