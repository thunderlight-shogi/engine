FROM golang:1.22.0 as build
COPY ./ /src/
RUN go build -o /app/app /src/cmd/thunderlight/thunderlight.go

FROM scratch
WORKDIR /app
COPY --from=build /app/app /app/app
CMD ["/app/app"]
