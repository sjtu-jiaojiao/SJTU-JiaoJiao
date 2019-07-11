FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-auth .
COPY --from=build /config.json .

CMD ["./api-auth","--registry=consul","--registry_address=$MICRO_REGISTRY_ADDRESS"]