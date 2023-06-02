FROM golang:1.18-alpine

WORKDIR /usr/src/lavka

COPY . .
RUN mkdir -p /usr/local/bin/
RUN go mod tidy
RUN go build -v -o /usr/local/bin/lavka

CMD ["/usr/local/bin/lavka"]