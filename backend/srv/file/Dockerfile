FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-file .
COPY --from=build /config.json .

CMD ["./srv-file","--registry=consul"]