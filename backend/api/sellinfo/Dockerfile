FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-sellInfo .
COPY --from=build /config.json .

CMD ["./api-sellInfo","--registry=consul"]