FROM golang:1.19.0-buster AS build

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

# COPY ./.env ./
# ถ้าใช้บรรทัดนี้จะbuild jenkins ไม่ผ่าน COPY failed: file not found in build context or
# excluded by .dockerignore: stat .env: file does not exist 

RUN go mod download

COPY ./ ./

RUN go build -o server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/server ./

# COPY .env ./
# ถ้าใช้บรรทัดนี้จะbuild jenkins ไม่ผ่าน COPY failed: file not found in build context or
# excluded by .dockerignore: stat .env: file does not exist 

CMD ["/server"]
