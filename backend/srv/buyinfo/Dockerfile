FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-buyinfo .
COPY --from=build /config.json .

CMD ["./srv-buyinfo","--registry=consul"]