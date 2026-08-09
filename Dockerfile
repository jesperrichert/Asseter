FROM golang:1.25.4 AS base
WORKDIR /app

FROM node:latest AS frontend
WORKDIR /Asseter.Frontend

COPY /Asseter.Frontend /Asseter.Frontend

RUN npm install
RUN npm run build

FROM base AS backend
WORKDIR /Asseter.Backend

COPY /Asseter.Backend /Asseter.Backend

RUN go build cmd/main.go 

FROM base AS release
WORKDIR /app

ENV FRONTEND_BUILD=/app/Asseter.Frontend/build

COPY --from=frontend /Asseter.Frontend/build/ /app/Asseter.Frontend/build/
COPY --from=backend /Asseter.Backend/main /app/Asseter.Backend/main

CMD [ "/app/Asseter.Backend/main" ]