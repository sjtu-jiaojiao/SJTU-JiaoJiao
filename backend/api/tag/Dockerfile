FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-tag .
COPY --from=build /config.json .

CMD ["./api-tag","--registry=consul"]