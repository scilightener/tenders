FROM golang:1.22-alpine as init

WORKDIR /app

COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify
COPY ./src .

FROM init as build
RUN CGO_ENABLED=0 go build -o migrator-tool ./cmd/migrator
RUN CGO_ENABLED=0 go build -o app ./cmd/app

FROM alpine
RUN echo '#!/bin/sh' > /entrypoint.sh && \
    echo '/opt/app/migrator-tool -migrations-path=/opt/app/migrations -direction=up' >> /entrypoint.sh && \
    echo '/opt/app/server' >> /entrypoint.sh && \
    chmod +x /entrypoint.sh

COPY --from=build /app/migrator-tool /opt/app/migrator-tool
COPY --from=init /app/migrations /opt/app/migrations

COPY --from=build /app/app /opt/app/server

EXPOSE 22313

ENTRYPOINT ["/entrypoint.sh"]