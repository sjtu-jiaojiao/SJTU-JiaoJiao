FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-file .
COPY --from=build /config.json .

CMD ["./api-file","--registry=consul"]