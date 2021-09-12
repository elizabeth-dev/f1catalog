FROM golang:1.14

RUN mkdir /api
COPY . /api
WORKDIR /api
RUN go build -o /app .

CMD ["/app"]