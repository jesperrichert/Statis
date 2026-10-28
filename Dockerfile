FROM golang:1.27.3 AS base
WORKDIR /app

FROM node:latest AS frontend
WORKDIR /Statis.Frontend

COPY /Statis.Frontend /Statis.Frontend

RUN npm install
RUN npm run build

FROM base AS backend
WORKDIR /Statis.Backend

COPY /Statis.Backend /Statis.Backend

RUN go build cmd/main.go 

FROM base AS release
WORKDIR /app

ENV FRONTEND_BUILD=/app/Statis.Frontend/build

COPY --from=frontend /Statis.Frontend/build/ /app/Statis.Frontend/build/
COPY --from=backend /Statis.Backend/main /app/Statis.Backend/main

CMD [ "/app/Statis.Backend/main" ]