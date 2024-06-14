FROM golang:1.22.1

EXPOSE 8080

WORKDIR /app

ADD go.mod .

COPY ./ ./

RUN go mod tidy
RUN go build -o ./build/app ./cmd/url-shortener/main.go

ENTRYPOINT ["./build/app"]