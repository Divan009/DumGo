FROM golang:alpine

# RUN apk add --no-cache git

# Set the Current Working Directory inside the container

RUN mkdir /app
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.

# COPY go.mod .
# COPY go.sum .

# downloading all mod and sum dependencies

COPY . .

RUN go mod download

ENV PORT 8080

CMD ["wget", "https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh"]
CMD ["chmod", "+x", "wait-for-it.sh"]
CMD ["./wait-for-it.sh", "db:5432", "--"]

CMD ["go", "run", "server.go"]
