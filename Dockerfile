FROM alpine:latest as build
WORKDIR /build
RUN wget -O go1.14.4 https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz \
    && export PATH=$PATH:/usr/local/go/bin
COPY . /app
WORKDIR /build
RUN go build *.go

FROM alpine:latest as release
WORKDIR /app
COPY --from=build /app main
ENTRYPOINT ["main"]