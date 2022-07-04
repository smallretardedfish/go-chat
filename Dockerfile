FROM golang:1.18 AS builder
WORKDIR /go/src/chat/
COPY . .
RUN go mod download
RUN go build -o app ./cmd/chat/

FROM scratch
#WORKDIR /
COPY --from=builder /go/src/chat/app ./app
ENTRYPOINT ["./app/228"]