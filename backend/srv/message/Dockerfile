FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-message .
COPY --from=build /config.json .

CMD ["./srv-message","--registry=consul"]