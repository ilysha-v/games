package auth

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/authboss.v1"
	_ "gopkg.in/authboss.v1/auth"
	_ "gopkg.in/authboss.v1/lock"
	_ "gopkg.in/authboss.v1/recover"
	_ "gopkg.in/authboss.v1/register"
	_ "gopkg.in/authboss.v1/remember"

	"github.com/aarondl/tpl"
	"github.com/gorilla/schema"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"

	"github.com/ilysha-v/games/backend/configuration"
)

var (
	Ab        = authboss.New()
	database  = NewMemStorer()
	viewsPath = configuration.GetViewsPath()
	templates = tpl.Must(tpl.Load(viewsPath+"/views", viewsPath+"/views/partials", "layout.html.tpl", funcs))
	schemaDec = schema.NewDecoder()
)

var funcs = template.FuncMap{
	"formatDate": func(date time.Time) string {
		return date.Format("2006/01/02 03:04pm")
	},
	"yield": func() string { return "" },
}

func setupAuthboss() {
	Ab.Storer = database
	// Ab.OAuth2Storer = database
	Ab.MountPath = "/api/auth"
	Ab.RootURL = `http://localhost:8080/api/`

	Ab.LayoutDataMaker = layoutData
	Ab.ViewsPath = "ab_views"
	Ab.RegisterOKPath = "/api/auth/register"
	Ab.AuthLoginOKPath = "/api/whoami"

	// Ab.OAuth2Providers = map[string]authboss.OAuth2Provider{
	// 	"google": authboss.OAuth2Provider{
	// 		OAuth2Config: &oauth2.Config{
	// 			ClientID:     ``,
	// 			ClientSecret: ``,
	// 			Scopes:       []string{`profile`, `email`},
	// 			Endpoint:     google.Endpoint,
	// 		},
	// 		Callback: aboauth.Google,
	// 	},
	// }

	b, err := ioutil.ReadFile(filepath.Join(viewsPath+"/views", "layout.html.tpl"))
	if err != nil {
		panic(err)
	}

	Ab.Layout = template.Must(template.New("layout").Funcs(funcs).Parse(string(b)))

	Ab.XSRFName = "csrf_token"
	Ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}

	Ab.CookieStoreMaker = NewCookieStorer
	Ab.SessionStoreMaker = NewSessionStorer

	Ab.Mailer = authboss.LogMailer(os.Stdout)

	Ab.Policies = []authboss.Validator{
		authboss.Rules{
			FieldName:       "email",
			Required:        true,
			AllowWhitespace: false,
		},
		authboss.Rules{
			FieldName:       "password",
			Required:        true,
			MinLength:       4,
			MaxLength:       8,
			AllowWhitespace: false,
		},
	}

	if err := Ab.Init(); err != nil {
		log.Fatal(err)
	}

	schemaDec.IgnoreUnknownKeys(true)
}

func SetupAuth() {
	// todo - generate unique
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`AbfYwmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==`)
	cookieStore = securecookie.New(cookieStoreKey, nil)
	sessionStore = sessions.NewCookieStore(sessionStoreKey)
	setupAuthboss()
}

func layoutData(w http.ResponseWriter, r *http.Request) authboss.HTMLData {
	currentUserName := ""
	userInter, err := Ab.CurrentUser(w, r)
	if userInter != nil && err == nil {
		currentUserName = userInter.(*User).Name
	}

	return authboss.HTMLData{
		"loggedin":               userInter != nil,
		"username":               "",
		authboss.FlashSuccessKey: Ab.FlashSuccess(w, r),
		authboss.FlashErrorKey:   Ab.FlashError(w, r),
		"current_user_name":      currentUserName,
	}
}
