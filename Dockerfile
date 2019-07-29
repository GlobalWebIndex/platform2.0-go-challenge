# Start from golang:1.12-alpine base image
FROM golang:1.12-alpine

# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh 

# adding alias
RUN echo "alias ll='ls -lah'" >> /root/.bashrc

# Set the Current Working Directory inside the container
WORKDIR /app/src

# install external dep
RUN go get -d github.com/gorilla/mux

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src /app/src

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# start using the web server 
ENTRYPOINT ./main

