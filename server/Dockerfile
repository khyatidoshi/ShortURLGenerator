FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8081

# Run
CMD ["/docker-gs-ping"]