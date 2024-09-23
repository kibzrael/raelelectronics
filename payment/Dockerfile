FROM golang:1.23.1-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV AIR_PORT=8004
EXPOSE 8004

CMD ["air"]
