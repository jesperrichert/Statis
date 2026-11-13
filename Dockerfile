FROM golang:1.27.4 AS base
WORKDIR /app

FROM oven/bun:slim AS frontend
WORKDIR /Statis.Frontend

COPY /Statis.Frontend/package.json /Statis.Frontend/package.json
RUN bun install

COPY /Statis.Frontend /Statis.Frontend

RUN bun run build

FROM base AS backend
WORKDIR /Statis.Backend

COPY /Statis.Backend/go.mod /Statis.Backend/go.mod
COPY /Statis.Backend/go.sum /Statis.Backend/go.sum

RUN go mod download

COPY /Statis.Backend /Statis.Backend

RUN go build cmd/main.go 

FROM base AS release
WORKDIR /app

ENV FRONTEND_BUILD=/app/Statis.Frontend/build

COPY --from=frontend /Statis.Frontend/build/ /app/Statis.Frontend/build/
COPY --from=backend /Statis.Backend/main /app/Statis.Backend/main

CMD [ "/app/Statis.Backend/main" ]