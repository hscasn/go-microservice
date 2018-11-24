# Build stage
FROM golang:latest AS build-env
ADD . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o start internal/app.go

# Final stage
FROM alpine 
WORKDIR /
RUN apk add curl && rm -rf /var/cache/apk
COPY --from=build-env /src/start /start
ENTRYPOINT ["/start"]
EXPOSE 8000
