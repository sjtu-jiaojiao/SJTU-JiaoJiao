FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-buyinfo .
COPY --from=build /config.json .

CMD ["./api-buyinfo","--registry=consul"]