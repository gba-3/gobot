FROM golang:1.16.0 as builder

RUN apt update && apt upgrade -y

WORKDIR /app
COPY . .
RUN make build

FROM debian:latest
RUN apt update && apt upgrade -y && mkdir -p /app

ENV SLACK_APP_TOKEN=${SLACK_APP_TOKEN}
ENV SLACK_BOT_TOKEN=${SLACK_APP_TOKEN}

WORKDIR /app
EXPOSE 9090
COPY --from=builder /app/gobot /app
