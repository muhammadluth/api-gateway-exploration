FROM golang:alpine AS build-env

LABEL maintainer="Muhammad Luthfi <muhammadluthfi059@gmail.com>"

ARG SERVICE_NAME=comment-backend

RUN mkdir /app
ADD . /app/

WORKDIR /app
RUN go build -mod=vendor -o ${SERVICE_NAME} .

FROM alpine
WORKDIR /app
COPY --from=build-env /app/${SERVICE_NAME}   /app/${SERVICE_NAME}

RUN mkdir -p logs

RUN apk upgrade libssl3 libcrypto3
RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta

ENTRYPOINT ["/app/comment-backend"]
