FROM golang:latest

WORKDIR /

COPY . .

RUN go build Dynamic

CMD [ "./Dynamic" ]