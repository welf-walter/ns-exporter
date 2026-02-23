FROM golang:1.25-alpine AS build

WORKDIR /app
COPY *.go go.* ./

RUN CGO_ENABLED=0 go build -o /ns-exporter .

FROM alpine

# see https://docs.github.com/de/packages/working-with-a-github-packages-registry/working-with-the-container-registry#labelling-container-images
LABEL org.opencontainers.image.source=https://github.com/welf-walter/ns-exporter
LABEL org.opencontainers.image.description="Nightscout exporter to InfluxDB"

COPY --from=build /ns-exporter /usr/local/bin/ns-exporter

# taken over from https://gitlab.com/DjPicLLC/docker-cron-image
COPY container/entrypoint.sh /entrypoint.sh
COPY container/crontab /crontab
RUN crontab /crontab

ENTRYPOINT /entrypoint.sh