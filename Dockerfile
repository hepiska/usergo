FROM golang:1.14-alpine as builder
WORKDIR /app
COPY . .
RUN apk add --no-cache --virtual git
RUN env GOOS=linux GOARCH=amd64 && go mod tidy && go build -o usergo .

FROM alpine
WORKDIR /app
COPY --from=builder /app/usergo /app/
COPY --from=builder /app/config.yml /app/
COPY --from=builder /app/.env /app/

CMD [ "./usergo"]
EXPOSE 4004