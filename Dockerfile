##Builder Image
FROM golang:1.14 as builder
ENV GO111MODULE=on
COPY . /message-go
WORKDIR /message-go
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application

#s Run Image
FROM scratch
COPY --from=builder /message-go/bin/application application
EXPOSE 9999
ENTRYPOINT ["./application"]