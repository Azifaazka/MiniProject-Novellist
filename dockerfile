# FROM golang:1.17-alpine

# WORKDIR /app

# Copy go.mod ./
# copy go.sum ./

# RUN go mod download

# COPY . .

# RUN go build -o /main

# EXPOSE 8080

# CMD ["./main"]

FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN go build -tags netgo -o main.app .


# ------------------------------------


FROM alpine:latest

WORKDIR /kemasan

COPY --from=builder /app/main.app .

CMD [ "/kemasan/main.app" ]