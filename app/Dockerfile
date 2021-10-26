FROM golang:1.17.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-api

EXPOSE 8080

CMD [ "/go-api" ]
