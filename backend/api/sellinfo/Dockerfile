FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-sellinfo .
COPY --from=build /config.json .

CMD ["./api-sellinfo","--registry=consul"]