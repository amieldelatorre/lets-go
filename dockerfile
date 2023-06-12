FROM golang:latest

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=1 GOOS=linux go build -v -o web.exe ./cmd/web
EXPOSE 4000

CMD ["./web"]