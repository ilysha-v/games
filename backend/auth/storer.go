package auth

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ilysha-v/authboss"
	"github.com/ilysha-v/games/backend/configuration"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID   int
	Name string

	// Auth
	Email    string
	Password string

	// OAuth2
	Oauth2Uid      string
	Oauth2Provider string
	Oauth2Token    string
	Oauth2Refresh  string
	Oauth2Expiry   time.Time

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Lock
	AttemptNumber int64
	AttemptTime   time.Time
	Locked        time.Time

	// Recover
	RecoverToken       string
	RecoverTokenExpiry time.Time

	// Remember is in another table
}

type MongoStorer struct {
	Users  map[string]User
	Tokens map[string][]string
}

func NewStorer() *MongoStorer {
	return &MongoStorer{}
}

func openConnection() (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial(configuration.GetDatabaseHost())
	if err != nil {
		panic(err)
	}

	collection := session.DB("games").C("users")
	return session, collection
}

func (s MongoStorer) Create(key string, attr authboss.Attributes) error {
	var user User
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	session, collection := openConnection()
	defer session.Close()

	err := collection.Insert(user)
	if err != nil {
		panic(err)
	}

	return nil
}

func (s MongoStorer) Put(key string, attr authboss.Attributes) error {
	return s.Create(key, attr)
}

func (s MongoStorer) Get(key string) (result interface{}, err error) {
	session, collection := openConnection()

	defer session.Close()

	var user User
	err = collection.Find(bson.M{"email": key}).One(&user)
	if err != nil {
		return nil, authboss.ErrUserNotFound
	}

	return &user, nil
}

func (s MongoStorer) PutOAuth(uid, provider string, attr authboss.Attributes) error {
	return s.Create(uid+provider, attr)
}

func (s MongoStorer) GetOAuth(uid, provider string) (result interface{}, err error) {
	user, ok := s.Users[uid+provider]
	if !ok {
		return nil, authboss.ErrUserNotFound
	}

	return &user, nil
}

func (s MongoStorer) AddToken(key, token string) error {
	fmt.Print("add called")
	s.Tokens[key] = append(s.Tokens[key], token)
	fmt.Println("AddToken")
	spew.Dump(s.Tokens)
	return nil
}

func (s MongoStorer) DelTokens(key string) error {
	delete(s.Tokens, key)
	fmt.Println("DelTokens")
	spew.Dump(s.Tokens)
	return nil
}

func (s MongoStorer) UseToken(givenKey, token string) error {
	fmt.Print("use called")
	toks, ok := s.Tokens[givenKey]
	if !ok {
		return authboss.ErrTokenNotFound
	}

	for i, tok := range toks {
		if tok == token {
			toks[i], toks[len(toks)-1] = toks[len(toks)-1], toks[i]
			s.Tokens[givenKey] = toks[:len(toks)-1]
			return nil
		}
	}

	return authboss.ErrTokenNotFound
}

func (s MongoStorer) ConfirmUser(tok string) (result interface{}, err error) {
	fmt.Println("==============", tok)

	for _, u := range s.Users {
		if u.ConfirmToken == tok {
			return &u, nil
		}
	}

	return nil, authboss.ErrUserNotFound
}

func (s MongoStorer) RecoverUser(rec string) (result interface{}, err error) {
	for _, u := range s.Users {
		if u.RecoverToken == rec {
			return &u, nil
		}
	}

	return nil, authboss.ErrUserNotFound
}
