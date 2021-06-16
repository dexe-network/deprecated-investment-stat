FROM golang:1.15-alpine AS gobuild


WORKDIR /build
COPY cmd ./cmd
COPY internal ./internal
COPY pkg ./pkg
COPY go.mod go.sum ./
RUN apk --no-cache add git mercurial ca-certificates
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=gobuild /build/.bin/app .
CMD ["./app"]
