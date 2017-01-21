FROM golang
ADD . /go/src/github.com/ilysha-v/games

#deps
RUN go get github.com/gorilla/mux
RUN go get github.com/juju/loggo
RUN go get gopkg.in/mgo.v2

RUN go install github.com/ilysha-v/games
ENTRYPOINT /go/bin/games
EXPOSE 8080
