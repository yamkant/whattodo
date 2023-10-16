FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY app ./
RUN go mod download

# Build
RUN CGO_ENABLED=1 GOOS=linux go build -o /docker-gs-ping
RUN ln -snf /usr/share/zoneinfo/Asia/Seoul /etc/localtime

# Optional:
EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]
