FROM golang:1.23rc2-alpine3.19

WORKDIR /app

COPY . .

RUN go build -o math

CMD [ "./math" ]