FROM golang:1.6
ENV VIEWS_PATH=/go/src/github.com/ilysha-v/games/backend/auth
COPY . /go/src/github.com/ilysha-v/games

#deps
RUN go get github.com/gorilla/mux
RUN go get github.com/juju/loggo
RUN go get gopkg.in/mgo.v2
RUN go get github.com/aarondl/tpl
RUN go get github.com/davecgh/go-spew/spew
RUN go get github.com/gorilla/schema
RUN go get github.com/gorilla/securecookie
RUN go get github.com/gorilla/sessions
RUN go get github.com/justinas/nosurf
RUN go get github.com/ilysha-v/authboss
RUN go get github.com/ilysha-v/authboss/auth
RUN go get github.com/ilysha-v/authboss/lock
RUN go get github.com/ilysha-v/authboss/recover
RUN go get github.com/ilysha-v/authboss/register
RUN go get github.com/ilysha-v/authboss/remember

RUN go install github.com/ilysha-v/games
ENTRYPOINT /go/bin/games
EXPOSE 8080
