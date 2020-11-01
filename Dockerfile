FROM golang:alpine3.12

WORKDIR /go/src/app
COPY . .
RUN go build -o main .
EXPOSE 3000
CMD [ "./main" ]
