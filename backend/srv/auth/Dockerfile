FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-auth .
COPY --from=build /config.json .

CMD ["./srv-auth","--registry=consul"]