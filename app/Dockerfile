FROM golang:1.16

WORKDIR /user/src/app

RUN export PATH="$PATH:$(go env GOPATH)/bin"

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# binary will be $(go env GOPATH)/bin/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN go build -o /main


EXPOSE 7001

CMD [ "/main" ]