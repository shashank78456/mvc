FROM golang:latest

WORKDIR /app

COPY . /app

RUN go mod download

RUN go build -o main ./cmd/

EXPOSE 3000

CMD [ "./main" ]
