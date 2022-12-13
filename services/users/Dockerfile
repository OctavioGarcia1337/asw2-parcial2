FROM golang:1.17

ADD . /users

WORKDIR /users
RUN go mod init users
RUN go mod tidy
RUN go build -o users .
RUN chmod +x /users

ENTRYPOINT ["/users/users"]
