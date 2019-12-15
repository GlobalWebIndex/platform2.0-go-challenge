FROM golang:1.13-alpine
RUN apk add --no-cache gcc musl-dev

ENV GO111MODULE=on

COPY src/ /app
WORKDIR /app

RUN go build

CMD ["./gwi-challenge"]

# docker build -t gwi-demo .
# docker run -p 8080:8080 gwi-demo