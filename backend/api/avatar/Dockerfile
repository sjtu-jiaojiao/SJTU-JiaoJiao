FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-avatar .
COPY --from=build /config.json .

CMD ["./api-avatar","--registry=consul"]