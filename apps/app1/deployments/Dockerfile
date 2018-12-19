# Build stage
FROM golang:latest AS build-env
ARG GOROOTREPO
ARG GOAPPREPO
ENV GOPATH=/go
ADD . ${GOPATH}/src/${GOROOTREPO}
WORKDIR ${GOPATH}/src/${GOROOTREPO}
RUN go get ./...
WORKDIR ${GOPATH}/src/${GOAPPREPO}
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o start ./internal/app.go
WORKDIR /
RUN mv ${GOPATH}/src/${GOAPPREPO}/start ./start

# Final stage
FROM alpine 
WORKDIR /
RUN apk add curl && rm -rf /var/cache/apk
COPY --from=build-env /start /start
ENTRYPOINT ["/start"]
EXPOSE 8000
