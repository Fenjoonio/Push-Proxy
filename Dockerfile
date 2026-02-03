FROM golang:1.24.2

ENV PORT=3000

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o /phoxy

EXPOSE 3000

ENTRYPOINT [ "/phoxy" ]
