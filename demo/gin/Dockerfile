FROM golang:1.16.8-alpine3.14 as build
ENV GIN_MODE=release
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /gin

FROM alpine:3.14
ENV GIN_MODE=release
WORKDIR /
COPY --from=build /gin /gin
COPY wait-for wait-for
ENTRYPOINT [ "./wait-for", "eureka:9000", "--", "/gin"]
