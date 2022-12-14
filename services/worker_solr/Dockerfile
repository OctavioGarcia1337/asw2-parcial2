FROM golang:1.18

ADD . /busqueda

WORKDIR /busqueda
RUN go mod init worker
RUN go mod tidy
RUN go build -o worker .
RUN chmod +x /busqueda

ENTRYPOINT ["/busqueda/worker"]
