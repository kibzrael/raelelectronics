FROM golang:1.23.1-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV AIR_PORT=8002
EXPOSE 8002

CMD ["air"]
