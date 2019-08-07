FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add --no-cache git gcc libc-dev
RUN go build -o main .
WORKDIR /app/apiTests
RUN go build -o ../test 
WORKDIR /app
EXPOSE 1323
CMD /app/main