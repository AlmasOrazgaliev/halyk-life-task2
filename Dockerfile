FROM golang:latest

RUN go version

ENV GOPATH=/
COPY ./ ./



RUN go mod download
RUN go build -o halyk-task ./main.go

CMD ["./halyk-task"]