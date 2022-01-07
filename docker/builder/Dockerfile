FROM golang:1.17
# RUN apk add --no-cache git gcc dumb-init
RUN mkdir /app
WORKDIR /app
# Add Tini
ENV TINI_VERSION v0.19.0
ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini

# Add supervisor
RUN git clone https://github.com/ochinchina/supervisord.git
WORKDIR /app/supervisord
RUN go generate
RUN GOOS=linux go build -tags release -a -ldflags "-linkmode external -extldflags -static" -o supervisord
RUN mv supervisord /app/supervisord-bin