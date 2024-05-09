FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -o myapp

ENV PORT=$PORT_APP
EXPOSE $PORT

CMD [ "./myapp" ]