# Creating GO-Lang enviornnemt as go-builder
FROM golang:1.21.5-alpine3.18 AS go-builder

# Creating app directory 
WORKDIR /app

# Copy Go Mod file in app directory
COPY go.mod ./

# Copy Go Sum file in app directory
COPY go.sum ./

# Downloading all external library needed
RUN go mod download -x

# Copy rest all files
COPY ./ ./

# Build Go appplication
RUN CGO_ENABLED=0 go build -o gqlgenExample ./server.go


# Creating an enviornment for go application to run
FROM alpine:3.18.5

# Copy Built executable file to the enviornment created
COPY --from=go-builder ["/app/gqlgenExample","/app/.env" ,"/root/"]

WORKDIR /root

# Update access permission for executable file 
RUN chmod +x gqlgenExample