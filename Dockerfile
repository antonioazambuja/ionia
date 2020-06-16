FROM golang:latest as builder
COPY . /app
WORKDIR /app
RUN go get -u github.com/gorilla/mux \
    && go get -u github.com/antonioazambuja/ionia \
    && go get -u github.com/testify/assert \\
    && CGO_ENABLED=0 GOOS=linux go build *.go

FROM alpine:latest as release
WORKDIR /app
COPY --from=builder /app/main app
CMD ["./app"]