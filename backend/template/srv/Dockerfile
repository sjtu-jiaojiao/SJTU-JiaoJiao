FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/srv-{{SERVICE_LNAME}} .
COPY --from=build /config.json .

CMD ["./srv-{{SERVICE_LNAME}}","--registry=consul"]