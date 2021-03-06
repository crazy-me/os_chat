FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o os_chat .
CMD ["/app/os_chat"]
