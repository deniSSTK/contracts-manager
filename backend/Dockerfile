FROM golang:1.25

RUN apt-get update && apt-get install -y gcc sqlite3 libsqlite3-dev bash

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000

CMD ["go", "run", "main.go"]