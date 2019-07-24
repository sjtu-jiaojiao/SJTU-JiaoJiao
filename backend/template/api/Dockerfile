FROM sjtujj:latest AS build
FROM scratch AS prod
COPY --from=build /build/api-{{SERVICE_LNAME}} .
COPY --from=build /config.json .

CMD ["./api-{{SERVICE_LNAME}}","--registry=consul"]