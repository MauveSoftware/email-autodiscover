FROM golang as builder
ADD . /go/src/
WORKDIR /go/src/cmd/autodiscover
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/email-autodiscover

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
COPY --from=builder /go/bin/email-autodiscover /app/email-autodiscover
EXPOSE 8080
VOLUME /config
ENTRYPOINT /app/email-autodiscover --http-port 8080 --config /config/config.yml
