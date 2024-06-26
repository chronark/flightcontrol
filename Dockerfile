# use the official Bun image
# see all versions at https://hub.docker.com/r/oven/bun/tags
FROM golang:1.22 as builder
WORKDIR /usr/src/app

COPY . .
RUN go mod tidy
RUN go build  -o main ./cmd/main.go

# run the app


FROM golang:1.22
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/main ./main

ARG PORT
EXPOSE ${PORT:-8080}

ENTRYPOINT [ "./main" ]