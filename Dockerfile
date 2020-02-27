FROM golang:1.13 as build
LABEL owner="jphuc96@gmail.com"
WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microst ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/microst .
EXPOSE 8080
CMD ["/app/microst"]
