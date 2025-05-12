FROM golang:1.23-alpine as builder

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine
WORKDIR /app/
COPY --from=builder /app/app .
COPY --from=builder /app/config/firebase-credentials.json /app/config/firebase-credentials.json
ENTRYPOINT ["./app"]
