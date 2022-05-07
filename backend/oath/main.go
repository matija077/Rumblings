package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/go-session/session"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var (
	dump   bool
	id     string
	secret string
	domain string
	port   int
)

const (
	POST string = "POST"
	GET         = "GET"
)

const (
	userKey   string = "User1"
	userValue        = "Data"
)

func init() {
	flag.BoolVar(&dump, "dump", false, "Dump requests and responses")
	flag.StringVar(&id, "id", "222222", "The client id being passed in")
	flag.StringVar(&secret, "secret", "22222222", "The client secret being passed in")
	flag.StringVar(&domain, "domain", "http://localhost:9094", "The domain of the redirect url")
	flag.IntVar(&port, "port", 9096, "the base port for the server")
	flag.Parse()
}

func main() {
	// creating a manager
	manager := manage.NewDefaultManager()
	//config for the manager for tokens
	cfg := manage.Config{
		AccessTokenExp:    time.Minute * 5,
		RefreshTokenExp:   time.Hour * 24,
		IsGenerateRefresh: true,
	}
	manager.SetAuthorizeCodeTokenCfg(&cfg)
	// creating a token storage
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// creating client store
	clientStore := store.NewClientStore()
	// creating one client
	client := &models.Client{
		ID:     id,
		Secret: secret,
		Domain: domain,
	}
	// seting one client
	clientStore.Set(id, client)

	//server creation
	srv := server.NewServer(server.NewConfig(), manager)

	srv.SetResponseErrorHandler(responseErrorHandler)
	srv.SetInternalErrorHandler(internalErrorHandler)
	srv.SetUserAuthorizationHandler(userAuthorizationHandler)
	srv.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/authorize", withAuthorizeHandler(srv))
	http.HandleFunc("/token", withTokenHandler(srv))

	log.Println("Server is running at  port.")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

// ROUTE HANDLERS
func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loginhandler")

	// create new session with current context. if error return 500
	sessionStore, err := session.Start(context.Background(), w, r)
	if err != nil {
		log.Fatalln("error creating a new session inside loginHandler")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// if get then redirect to login page
	if r.Method != POST {
		outputHTML(w, r, "static/login.html")
		return
	}

	// if post then save user's data to session store. this is currently hardcoded. also redirect to
	// auth page and handle auth
	// TODO
	sessionStore.Set(userKey, userValue)
	sessionStore.Save()

	w.Header().Set("Location", "/auth")
	w.WriteHeader(http.StatusFound)

}
func authHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("authHandler")

	sessionStore, err := session.Start(context.Background(), w, r)
	if err != nil {
		log.Fatalln("error when starting session in authHandler")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, ok := sessionStore.Get(userKey); !ok {
		log.Println("user not found in session,. redirect to login from authHandler")
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	if r.Method == POST {
		var form url.Values
		if v, ok := sessionStore.Get("eturnUri"); ok {
			form, ok = v.(url.Values)
			if !ok {
				log.Fatalln("panic type assertion from url.Values in authHandler")
			}
		}

		url := new(url.URL)
		url.Path = "/authorize"
		url.RawQuery = form.Encode()
		w.Header().Set("Location", url.String())
		w.WriteHeader(http.StatusFound)
		sessionStore.Delete("Form")


        if v, ok := store.Get(userKey); ok {
            store.Set("UserID", v)
        }
        store.Save()

        return
    }
    outputHTML(w, r, "static/auth.html")
	}

}
func withAuthorizeHandler(srv *server.Server) func(w http.ResponseWriter, r *http.Request) {
	var authorizeHandler = func(w http.ResponseWriter, r *http.Request) {
		log.Println("authorizeHandler")

		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			log.Fatalln("error inside authorizeHandler")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return authorizeHandler
}
func withTokenHandler(srv *server.Server) func(w http.ResponseWriter, r *http.Request) {
	var tokenHandler = func(w http.ResponseWriter, r *http.Request) {
		log.Println("tokenHandler")

		err := srv.HandleTokenRequest(w, r)
		if err != nil {
			log.Fatalln("error inside tokenHandler")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	return tokenHandler
}

// OAUTH SERVER HANDLERS
func responseErrorHandler(re *errors.Response) {
	log.Println("Response Error: ", re.Error.Error())
}
func internalErrorHandler(err error) (re *errors.Response) {
	log.Println("Internal Error: ", err.Error())
	return re
}
func userAuthorizationHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	sessionStore, err := session.Start(context.Background(), w, r)
	if err != nil {
		log.Fatalln("error when starting session in userAuthorizationHandler")
		return userID, err
	}

	// TODO hardocded
	uid, ok := sessionStore.Get(userKey)
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}

		sessionStore.Set("ReturnUri", r.Form)
		sessionStore.Save()

		log.Println("user not found in session. redirect to login from userAuthorizationHandler")
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return userID, err
	}

	userID, ok = uid.(string)
	if !ok {
		log.Fatalln("panic when type assertion from uid to userID in userAuthorizationHandler")
	}
	sessionStore.Delete(userKey)
	sessionStore.Save()

	return userID, err
}
func passwordAuthorizationHandler(username, password string) (userID string, err error) {
	return userID, err
}

// HELPERS
func outputHTML(w http.ResponseWriter, r *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("error opening the file in outputHTML")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln("error opening the file info in outputHTML")
	}

	http.ServeContent(w, r, file.Name(), fileInfo.ModTime(), file)
}
