# build Stage
FROM golang:alpine AS build

# installing build dependencies
RUN apk add --no-cache gcc musl-dev make

WORKDIR /app

COPY . .

RUN make

FROM alpine:latest

RUN mkdir /app

COPY --from=build /app/cmd/sn-go-api /app/

WORKDIR /app

COPY . /app/

EXPOSE 8080

CMD ["./sn-go-api"]