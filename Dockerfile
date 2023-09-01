FROM golang:1.21-alpine

ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o avito-backend-task ./cmd/app/main.go

CMD ["./avito-backend-task"]