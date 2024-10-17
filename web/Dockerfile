FROM golang:1.23.1-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ARG PORT

ENV PORT=${PORT}
EXPOSE ${PORT}

CMD ["air", "-c", ".air.toml"]
