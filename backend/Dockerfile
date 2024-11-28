FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /blog-scraper-qa ./cmd/main

EXPOSE 8000

CMD [ "/blog-scraper-qa" ]