# syntax=docker/dockerfile:1

FROM --platform=linux/amd64 golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download
# 6719b760ab219e96ff30f1ff
# 6719b77aab219e96ff30f200
# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=1 GOOS=linux go build -o /feed-complex-service

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8089

# Run
CMD ["/feed-complex-service"]