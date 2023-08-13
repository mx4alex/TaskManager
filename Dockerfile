FROM golang:1.20

WORKDIR /usr/src/app

COPY ./ ./

# install postgres
RUN apt-get update
RUN apt-get -y install postgresql-client

# wait postgres
RUN chmod +x entrypoint.sh

# build go app
RUN go mod download
RUN go build -o taskmanager ./cmd/app/main.go

CMD ["./taskmanager"]