FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-sellinfo .
COPY --from=build /config.json .

CMD ["./srv-sellinfo","--registry=consul"]