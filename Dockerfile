FROM golang:1.19.0-buster AS build

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./
COPY ./.env ./

RUN go mod download

COPY ./ ./

RUN go build -o server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/server ./
COPY .env ./

CMD ["/server"]
