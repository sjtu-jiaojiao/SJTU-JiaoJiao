FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-avatar .
COPY --from=build /config.json .

CMD ["./srv-avatar","--registry=consul"]