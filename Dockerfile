FROM golang:1.20

WORKDIR /usr/src/app

# install psql
RUN apt-get update \
    && apt-get -y install postgresql-client

# install dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

# build app
RUN chmod +x entrypoint.sh \
    && go build -o taskmanager ./cmd/app/main.go

CMD ["./taskmanager"]