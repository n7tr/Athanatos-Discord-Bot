# For deployment on railway.app
FROM golang:latest

WORKDIR /Dynamic

COPY . .

RUN go build Dynamic

CMD [ "./Dynamic" ]
