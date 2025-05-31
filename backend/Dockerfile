FROM golang:1.21-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

EXPOSE 3000
CMD ["go", "run", "main.go"] 