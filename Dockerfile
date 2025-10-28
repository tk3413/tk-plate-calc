# build
FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app .

# runtime
FROM alpine:3.18
RUN adduser -D appuser
COPY --from=build /app /app
USER appuser
EXPOSE 8080
ENTRYPOINT ["/app"]