FROM golang:1.18

ADD . /busqueda

WORKDIR /busqueda
RUN go mod init wesolr
RUN go mod tidy
RUN go build -o wesolr .
RUN chmod +x /busqueda

ENTRYPOINT ["/busqueda/wesolr"]
