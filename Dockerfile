FROM golang:1.20.7-alpine3.18

WORKDIR /app

COPY go.mod ./
COPY db ./db
COPY *.go ./
COPY static ./static

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /server


CMD ["/server"]