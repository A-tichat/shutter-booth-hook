FROM golang:1.22 AS dependencies
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
COPY --from=dependencies /go/pkg /go/pkg
RUN make build

FROM alpine:3.20
ARG PORT 3000
EXPOSE ${PORT}
WORKDIR /app
COPY --from=builder /app/build ./
CMD ["./app"]