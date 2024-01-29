FROM golang:1.21.5
WORKDIR /app
COPY go.mod go.sum ./
EXPOSE 5052
RUN go mod download
COPY . ./
RUN go build -o geoip-service ./cmd
CMD ["./geoip-service"]