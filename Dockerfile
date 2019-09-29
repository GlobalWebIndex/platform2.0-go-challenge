FROM golang:alpine 

# Add Maintainer info
LABEL maintainer="Elias Krontiris <teliaz@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 



RUN go build -o gwiapi . 
EXPOSE 8080
CMD ["/app/gwiapi"]
