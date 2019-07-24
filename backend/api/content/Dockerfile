FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-content .
COPY --from=build /config.json .

CMD ["./api-content","--registry=consul"]