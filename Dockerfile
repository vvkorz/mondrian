# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /webapp
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY app/ ./app
COPY mondrian/ ./mondrian
RUN ls -la ./
RUN ls -la ./app
RUN ls -la ./mondrian
WORKDIR /webapp/app
RUN ls -la ./
#RUN go env -w GO111MODULE=off
RUN go build -o ./mondrianapp
RUN ls -la ./
EXPOSE 8080

CMD [ "/webapp/app/mondrianapp" ]