# Requires Docker 17.05.0 (https://github.com/samsung-cnct/issues/issues/55)
FROM grpc/go:1.0 as build

FROM gcr.io/distroless/base
