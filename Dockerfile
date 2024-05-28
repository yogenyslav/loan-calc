FROM golang:1.21.0-alpine as builder
WORKDIR /builder

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./main ./cmd/server/main.go

FROM scratch as prod
WORKDIR /app

COPY config/config.yml ./config/
COPY --from=builder /builder/main .

CMD [ "./main" ]