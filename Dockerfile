# dev, builder
FROM golang:1.14 AS golang
WORKDIR /work/yatter-backend-go
ENV GO111MODULE=on

# dev
FROM golang as dev
RUN curl -fLo /usr/local/bin/air https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air \
  && chmod +x /usr/local/bin/air

# builder
FROM golang AS builder
COPY ./ ./
RUN make prepare build-linux

# release
FROM alpine AS app
RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=builder /work/yatter-backend-go/build/yatter-backend-go-linux-amd64 /usr/local/bin/yatter-backend-go
EXPOSE 8080
ENTRYPOINT ["yatter-backend-go"]
