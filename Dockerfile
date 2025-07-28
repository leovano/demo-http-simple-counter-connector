FROM golang:1 as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM gcr.io/distroless/static-debian12

COPY --from=build /app/main /

EXPOSE 12345
CMD ["/main"]
