FROM golang:1.18-bullseye

RUN go mod download 

EXPOSE 8080
ENV APP_HOME /go/src/mathapp
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"
CMD ["go", "run" "main.go"]