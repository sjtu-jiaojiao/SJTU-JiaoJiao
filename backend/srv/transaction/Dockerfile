FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-transaction .
COPY --from=build /config.json .

CMD ["./srv-transaction","--registry=consul"]