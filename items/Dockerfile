FROM golang:1.18

ADD . /backend

WORKDIR /backend
RUN go mod init items
RUN go mod tidy
RUN go build -o items .
RUN chmod +x /backend

ENTRYPOINT ["/backend/items"]
