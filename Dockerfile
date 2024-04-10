FROM golang:1.22.2-alpine3.19 AS BuildStage

WORKDIR /app
COPY . .

RUN go mod vendor
RUN go build -o /api cmd/api/api.go

FROM alpine:latest

WORKDIR /
COPY --from=BuildStage /api /api

EXPOSE 8080
ENTRYPOINT [ "/api" ]