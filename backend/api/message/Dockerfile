FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-message .
COPY --from=build /config.json .

CMD ["./api-message","--registry=consul"]