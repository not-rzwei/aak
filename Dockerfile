FROM golang:1.16-alpine

WORKDIR /app

COPY . ./
RUN go mod download

COPY cmd/ pkg/ ./

RUN go build -o /aak ./cmd/aak
EXPOSE 8080

ENTRYPOINT ["/aak"]
