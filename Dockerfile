FROM golang:1.16-alpine3.13 as build-env
RUN apk add --no-cache git gcc
RUN mkdir /app
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o bin/todo src/main.go

FROM gcr.io/distroless/static
COPY --from=build-env /app/bin/todo .
EXPOSE 3030/tcp
USER 1001
CMD ["./todo"]