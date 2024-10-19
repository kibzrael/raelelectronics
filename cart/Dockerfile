FROM golang:1.23.2-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV AIR_PORT=8003
EXPOSE 8003

CMD ["air"]