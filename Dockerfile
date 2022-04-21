FROM golang:1.18.1-alpine3.15 AS build
RUN mkdir /app 
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

FROM alpine:3.15 AS production
COPY --from=build /app/app /app/routes.json ./
CMD ["./app"]