FROM golang:1.19.3-alpine
WORKDIR /usr/src/app

COPY . .

RUN go get
RUN go install github.com/cosmtrek/air@v1.40.2