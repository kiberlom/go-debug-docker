FROM golang:1.17 as debug

RUN apt-get update && apt-get upgrade && \
    apt-get install -y \
        git \
        dpkg \
        gcc \
        musl-dev

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# эта штука и запускает дебаг 
RUN go get github.com/go-delve/delve/cmd/dlv@latest

RUN mkdir -p /app/src
WORKDIR /app/src
COPY ./app/src .
RUN go mod tidy

# only for build
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

COPY ./app/dlv.sh /
RUN chmod +x /dlv.sh
ENTRYPOINT [ "/dlv.sh" ]

# прод
FROM alpine:latest as build
COPY --from=debug /app/main .
EXPOSE 5045
CMD ["./main"]