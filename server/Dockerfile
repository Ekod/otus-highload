FROM golang:1.18-alpine AS build

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
COPY . .

RUN go build -o server ./cmd/app

FROM alpine

COPY --from=build /app/server .

EXPOSE 5000

CMD [ "./server" ]