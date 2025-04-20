FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o phoxy

CMD ["./phoxy"]

EXPOSE 8080
