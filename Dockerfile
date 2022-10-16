# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /webapp
COPY go.mod go.sum ./
RUN go mod download
COPY app/ ./app
COPY mondrian/ ./mondrian
WORKDIR /webapp/app
RUN go build -o ./mondrianapp
EXPOSE 8000

CMD [ "/webapp/app/mondrianapp" ]