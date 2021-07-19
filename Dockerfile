FROM golang:1.16.0 as builder

RUN apt update && apt upgrade -y

WORKDIR /app