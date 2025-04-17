FROM golang:1.24.2-alpine AS build

WORKDIR /go/src/app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go install -v ./cmd/...

FROM alpine:latest as sp_predictions_bot

WORKDIR /app

COPY --from=build /go/bin/main sp_predictions_bot

CMD ["./sp_predictions_bot"]
