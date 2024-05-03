FROM --platform=$BUILDPLATFORM golang:1.22-alpine as builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o bin/starling_exporter cmd/starling_exporter.go

FROM scratch

COPY --from=builder /app/bin/starling_exporter /starling_exporter
COPY --from=builder /etc/ssl /etc/ssl

ENTRYPOINT ["/starling_exporter"]
