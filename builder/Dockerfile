FROM golang:alpine AS builder

# Specifies the working dir in the new container
WORKDIR /go/src/go-boilerplate

ADD ./ .

# Runs the command in WORKDIR in the image
RUN go build -mod=vendor -o "go-boilerplate"

FROM alpine:latest
WORKDIR /bin/go-boilerplate

COPY --from=builder /go/src/go-boilerplate/go-boilerplate .
RUN mkdir config
COPY --from=builder /go/src/go-boilerplate/config/tier ./config/tier

CMD ["./go-boilerplate"]
EXPOSE 80 5001
