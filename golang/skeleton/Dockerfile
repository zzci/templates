FROM golang:1.19.3-alpine3.16 AS build

RUN apk add --no-cache ca-certificates build-base git;

WORKDIR /code

COPY . /code/

RUN  make

FROM alpine:3.16

EXPOSE 3000

WORKDIR /app/key

COPY --from=build /code/build/keyd /app/key

CMD ["/app/key/keyd"]
