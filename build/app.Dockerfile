FROM golang:1.22.0-alpine3.19 as build
WORKDIR /src
COPY go.mod go.sum /src/
RUN go mod download && go mod verify
COPY cmd /src/cmd
COPY internal /src/internal
COPY pkg /src/pkg
RUN go build -o /app/app /src/cmd/thunderlight/thunderlight.go

FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/app /app/app
COPY website /app/website
EXPOSE 88
CMD ["/app/app"]
