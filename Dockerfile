FROM golang:1.18.1-alpine
EXPOSE 8080
WORKDIR "/app"
COPY . /app
WORKDIR /app
CMD ["go", "run", "main.go"]