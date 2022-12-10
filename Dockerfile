FROM golang:1.18.3-alpine3.16

WORKDIR /app

COPY ./ ./

RUN go build -o /bin/app ./

ENTRYPOINT ["app"]