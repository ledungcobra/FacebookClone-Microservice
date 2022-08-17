FROM golang:1.18-alpine as build
COPY . /app
WORKDIR /app
RUN go build -o gateway-go ./cmd

FROM alpine:3.6
WORKDIR /app
COPY --from=build /app/gateway-go ./
ENTRYPOINT [ "./gateway-go" ]
