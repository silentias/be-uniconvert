FROM golang:1.21-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o be-uniconvert ./cmd/main.go

FROM alpine:3.18

RUN apk add --no-cache ffmpeg bash ca-certificates

WORKDIR /app

COPY --from=builder /app/be-uniconvert .
что 
RUN mkdir -p /app/uploads

EXPOSE 8000

CMD ["./be-uniconvert"]