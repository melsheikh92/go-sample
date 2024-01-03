FROM golang:1.21.5-alpine3.18 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM gcr.io/distroless/static-debian11
COPY --from=builder /app .
EXPOSE 80
CMD ["./main"]

