# build Stage
FROM golang:alpine AS build

# installing build dependencies
RUN apk add --no-cache gcc musl-dev sqlite-libs

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 CGO_CFLAGS="-D_LARGEFILE64_SOURCE" go build -o ./build ./cmd/companies_service

FROM alpine:latest

RUN apk add --no-cache sqlite-libs musl-dev

RUN mkdir /app

COPY --from=build /app/build/companies_service /app/

WORKDIR /app

COPY . /app/

EXPOSE 8080

CMD ["./companies_service"]