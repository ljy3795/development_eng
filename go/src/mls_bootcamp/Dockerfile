# docker build -t mls-bootcamp .
FROM golang:alpine

# Set env vars.
ENV SEOUL_API_KEY="6f536d76626c6a793534526f6f7a57"
ENV MARIADB_USER="colin"
ENV MARIADB_PASS="aa"
ENV MARIADB_ADDR="172.17.0.2"
ENV MARIADB_PORT="3306"

# Make App
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .

# Run App
CMD ["/app/main"]

# docker run -d -p 3333:8080 --rm --name go-mls-bootcamp mls-bootcamp:latest



# FROM ubuntu:latest

# # Set env vars.
# ENV SEOUL_API_KEY="6f536d76626c6a793534526f6f7a57"
# ENV MARIADB_USER="colin"
# ENV MARIADB_PASS="aa"

# docker run -d -i -t --name ubuntu_test ubuntu_latest:latest