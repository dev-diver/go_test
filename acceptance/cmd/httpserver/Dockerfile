FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr acceptance/cmd/httpserver/*.go

EXPOSE 8080
CMD [ "./svr" ]