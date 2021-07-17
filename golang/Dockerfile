FROM golang:1.16-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

COPY . .

RUN go mod download
RUN go build cmd/main.go

WORKDIR /dist

RUN cp /build/main .

FROM alpine

COPY --from=builder /dist/main /app/main

WORKDIR /app

EXPOSE 3000

ENTRYPOINT ["./main"]

