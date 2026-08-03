FROM golang:1.26.5 AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /todo-build ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder ./todo-build ./todo-server

COPY --from=builder ./build/web ./web

COPY --from=builder ./build/.env .

CMD [ "./todo-server" ]
