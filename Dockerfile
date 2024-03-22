# syntax=docker/dockerfile:1

FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY /app/go.mod /app/go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY /app ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# 
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 3000

# Run
CMD ["/server"]