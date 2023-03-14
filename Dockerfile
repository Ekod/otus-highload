FROM golang:1.18-alpine AS build

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
COPY . .

RUN go build -o server ./cmd/app

RUN apk add --no-cache make

RUN wget https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz \
    && tar -xvf migrate.linux-amd64.tar.gz                      \
    && rm migrate.linux-amd64.tar.gz


RUN make migrate

FROM gcr.io/distroless/base-debian10

COPY --from=build /app/server .

EXPOSE 5000

CMD [ "./server" ]