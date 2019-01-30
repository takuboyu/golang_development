
FROM golang:1.11.1
WORKDIR /go/src/mygoapp
COPY . .

#Go Environment variable
ENV GO111MODULE=off

#Debag
RUN go get github.com/derekparker/delve/cmd/dlv

#install Go web framework
RUN go get -u github.com/labstack/echo

#install "Fresh" restart tool
RUN go get github.com/pilu/fresh

CMD ["dlv", "--listen=:8000", "--headless=true", "--api-version=2", "exec", "/server"]
CMD ["fresh"]

