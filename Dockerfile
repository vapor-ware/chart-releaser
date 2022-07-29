FROM scratch

LABEL org.label-schema.schema-version="1.0" \
      org.label-schema.name="chartreleaser/chart-releaser" \
      org.label-schema.vcs-url="https://github.com/vapor-ware/chart-releaser" \
      org.label-schema.vendor="Vapor IO"

ADD chart-releaser /bin

WORKDIR /release

ENTRYPOINT ["/bin/chart-releaser"]
