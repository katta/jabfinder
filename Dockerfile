FROM golang:1.14-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /tmp/jabfinder
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build


FROM alpine:3.9
RUN apk add ca-certificates
WORKDIR /app
RUN mkdir -p data
COPY --from=build_base /tmp/jabfinder/templates /app/templates
COPY --from=build_base /tmp/jabfinder/.jabfinder.yaml /app/
COPY --from=build_base /tmp/jabfinder/jabfinder /app/jabfinder

CMD ["./jabfinder","check","-d","294","-a","45","-n"]