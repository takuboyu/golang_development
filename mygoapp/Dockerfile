
FROM golang:1.11.1
WORKDIR /go/src/golang_development/mygoapp
COPY . .

#Go Environment variable
ENV GO111MODULE=on

#install Go web framework
RUN go get -u github.com/labstack/echo

#install "Fresh" restart tool
RUN go get github.com/pilu/fresh

CMD ["fresh"]

