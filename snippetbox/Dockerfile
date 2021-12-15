FROM golang:latest

WORKDIR /app

COPY . .

RUN go build ./cmd/web -o /snippetbox

CMD ["/snippetbox"]
