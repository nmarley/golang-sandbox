FROM golang:alpine
WORKDIR /app
COPY . /app
CMD ["/usr/local/go/bin/go", "run", "readblock.go"]
