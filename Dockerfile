# build Stage
FROM golang:alpine AS build

# installing build dependencies
RUN apk add --no-cache gcc musl-dev make

WORKDIR /app

COPY . .

RUN make build

FROM alpine:latest

RUN mkdir /app

COPY --from=build /app/build/companies_service /app/
COPY --from=build /app/build/tg_bot /app/

WORKDIR /app

COPY . /app/

EXPOSE 8080

CMD ["./companies_service"]