FROM golang:1.23-alpine AS builder



WORKDIR /go/src/github.com/unkeyed/unkey/apps/apprunner
COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

COPY . .
RUN go build -o bin/apprunner ./main.go


FROM golang:1.23-alpine
WORKDIR  /usr/local/bin
COPY --from=builder /go/src/github.com/unkeyed/unkey/apps/apprunner/bin/apprunner .

CMD [ "/usr/local/bin/apprunner"]
