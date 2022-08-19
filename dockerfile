FROM golang:1.18

WORKDIR /usr/src/app

RUN apt update
RUN apt install sqlite3

RUN go mod download github.com/gorilla/mux@v1.8.0
RUN go mod download github.com/mattn/go-sqlite3@latest

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app

CMD ["app"]