FROM golang:1.16-alpine3.13 as build-env
RUN apk add --no-cache git gcc make
RUN mkdir /app
WORKDIR /app
ADD . .
RUN make build

FROM gcr.io/distroless/static
COPY --from=build-env /app/bin/todo .
EXPOSE 3030/tcp
USER 1001
CMD ["./todo"]