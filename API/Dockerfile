FROM golang:1.14

WORKDIR /go/src/app
COPY ./API .

RUN go mod download
RUN go install app.go


CMD ["app"]