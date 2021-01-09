FROM golang
RUN mkdir /build
WORKDIR /build
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build .

FROM alpine:latest
EXPOSE 8080
EXPOSE 9090
COPY --from=0 /build/go-prom-poc go-prom-poc
CMD ./go-prom-poc