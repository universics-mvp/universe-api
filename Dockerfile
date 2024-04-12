FROM golang:1.22-alpine as BUILDER

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o main main.go

FROM alpine:latest as RUNNER

WORKDIR /app

COPY --from=BUILDER /app/main /app/main

CMD ["/app/main"]