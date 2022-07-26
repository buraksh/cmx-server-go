FROM golang:1.18.3-alpine3.15

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o server

EXPOSE 80

CMD ["./server"]