FROM golang:1.18.9-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env.development ./
COPY .env.production ./

RUN go build -o /run

EXPOSE 8080

RUN chmod +x /run
CMD ["/run"]