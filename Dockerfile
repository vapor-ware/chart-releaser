# syntax=docker/dockerfile:1
# This docker file is mean to be used with goreleaser
# To build an image, do `goreleaser release --rm-dist --skip-publish`
FROM alpine/git:v2.36.2

LABEL org.label-schema.schema-version="1.0" \
      org.label-schema.name="chartreleaser/chart-releaser" \
      org.label-schema.vcs-url="https://github.com/vapor-ware/chart-releaser" \
      org.label-schema.vendor="Vapor IO"
COPY chart-releaser /usr/local/bin/chart-releaser

ENV WORKDIR /data
WORKDIR /data

ENTRYPOINT ["chart-releaser"]
