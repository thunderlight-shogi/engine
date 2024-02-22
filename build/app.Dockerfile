FROM golang:1.22.0-alpine as build
WORKDIR /src
COPY cmd /src/cmd
COPY internal /src/internal
COPY website /src/website
COPY go.mod go.sum /src/
RUN go build -o /app/app /src/cmd/thunderlight/thunderlight.go

FROM alpine:3.19.1
WORKDIR /app
COPY --from=build /app/app /app/app
EXPOSE 80
CMD ["/app/app"]
