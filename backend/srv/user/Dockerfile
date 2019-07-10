FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-user .
COPY --from=build /config.json .

CMD ["./srv-user","--registry=consul"]