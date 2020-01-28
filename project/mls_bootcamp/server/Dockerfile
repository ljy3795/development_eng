# docker build -t mls-bootcamp-server .
FROM golang:alpine

# Make App
WORKDIR /app
# RUN mkdir /app
COPY . .
RUN go build -o main .

# Run App
CMD ["/app/main"]