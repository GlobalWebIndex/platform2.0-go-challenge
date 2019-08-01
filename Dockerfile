# build stage
FROM golang:1.12.6-alpine3.9 AS builder
ADD . /src
RUN cd /src && go build -o gwichallenge

# final stage
FROM alpine:3.9
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/gwichallenge /app/
ENTRYPOINT ./gwichallenge