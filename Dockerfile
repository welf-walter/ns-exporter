FROM golang:1.25-alpine AS build

WORKDIR /app
COPY *.go go.* ./

RUN CGO_ENABLED=0 go build -o /ns-exporter .

FROM djpic/cron:standard

# see https://docs.github.com/de/packages/working-with-a-github-packages-registry/working-with-the-container-registry#labelling-container-images
LABEL org.opencontainers.image.source=https://github.com/welf-walter/ns-exporter
LABEL org.opencontainers.image.description="Nightscout exporter to InfluxDB"

COPY --from=build /ns-exporter /etc/periodic/1min/ns-exporter

RUN chmod 755 /etc/periodic/1min/ns-exporter
