# FROM vedaga/scjty8s4erci:builder-go-1.17-alpine3.15 as build-env
FROM vedaga/scjty8s4erci:builder-go-1.17-alpine3.13 as build-env
# FROM vedaga/scjty8s4erci:builder-supervisor as build-env
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o todo app.go

FROM gcr.io/distroless/static as prod
# FROM golang:1.17 as prod
# FROM krallin/ubuntu-tini:trusty as prod
ARG SERVER_PORT=3030
ENV SERVER_PORT=${SERVER_PORT}
COPY --from=build-env /app/todo /todo
EXPOSE ${SERVER_PORT}
USER 1001
ENV GOMAXPROCS=4
CMD [ "/todo" ]
