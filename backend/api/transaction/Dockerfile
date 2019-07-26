FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-transaction .
COPY --from=build /config.json .

CMD ["./api-transaction","--registry=consul"]