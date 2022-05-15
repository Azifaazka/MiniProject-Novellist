FROM golang:1.17-alpine

WORKDIR /app

Copy go.mod ./
copy go.sum ./

RUN go mod download

COPY . . 

RUN go build -o /dist

EXPOSE 3222

CMD ["/dist"]