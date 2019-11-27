FROM golang:1.13-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make
WORKDIR /app
COPY wait-for /usr/bin/wait-for
RUN chmod +x /usr/bin/wait-for
COPY . .