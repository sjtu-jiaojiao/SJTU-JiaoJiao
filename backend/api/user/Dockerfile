FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-user .
COPY --from=build /config.json .

CMD ["./api-user","--registry=consul"]