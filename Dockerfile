FROM golang:1.19.4

WORKDIR /app

COPY ./ ./

RUN go build -o /bin/app main.go

ENTRYPOINT ["app"]