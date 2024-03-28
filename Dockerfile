FROM golang:1.22

COPY . /app
WORKDIR /app

RUN go mod tidy
RUN go build -o /app/server /app/main.go

# expose port 50051
EXPOSE 50051

CMD ["/app/server"]