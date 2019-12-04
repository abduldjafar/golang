package app

import (
	"backend-code/auth"
	"backend-code/config"
	"backend-code/controller/Account"
	"backend-code/controller/Movies"
	"backend-code/controller/Optional"
	"backend-code/model"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router     *mux.Router
	SubRouter  *mux.Router
	TbAccounts *gorm.DB
	TbMovies   *gorm.DB
}

func (a *App) Initialize() {
	baseConfig := &config.Configuration{}
	config.GetConfig(baseConfig)

	db, err := gorm.Open("postgres", "host="+baseConfig.Postgres.Url+" port="+baseConfig.Postgres.Port+""+
		" user="+baseConfig.Postgres.User+" dbname="+baseConfig.Postgres.Db+" password="+baseConfig.Postgres.Password+
		" sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	a.TbAccounts = model.DBMigrationAccount(db, &model.Accounts{})
	a.TbMovies = model.DBMigrationAccount(db, &model.Movies{})
	a.Router = mux.NewRouter()
	a.SubRouter = a.Router.PathPrefix("/auth").Subrouter()
	a.SubRouter.Use(auth.JwtVerify)
	a.setRouters()
}

func (a *App) setRouters() {
	// Routing for handling the projects

	a.Post("/v1/account/register", a.createNewUser)
	a.Post("/v1/account/login", a.LoginAccount)
	a.PostAuth("/v1/account/delete", a.DeleteAccount)
	a.PostAuth("/v1/account/full", a.ListAccounts)
	a.PutAuth("/v1/account", a.UpdateAccount)
	a.PostQuery("/v1/movies", a.ListMovies)
	a.GetQuery("/v1/refresh", a.Optional)

}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for GET method with JWT
func (a *App) GetAuth(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("GET")
}

//get method with query
func (a *App) PostQuery(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("POST").Queries("count", "{count}")
}

func (a *App) GetQuery(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("GET").Queries("email", "{email}")
}

// Wrap the router for PUT method with JWT
func (a *App) PutAuth(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for POS method with JWT
func (a *App) PostAuth(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.SubRouter.HandleFunc(path, f).Methods("POST")
}

func (a *App) createNewUser(w http.ResponseWriter, r *http.Request) {
	Account.CreateNewUser(a.TbAccounts, w, r)
}

func (a *App) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	Account.UpdateAccount(a.TbAccounts, w, r)
}

func (a *App) LoginAccount(w http.ResponseWriter, r *http.Request) {
	Account.Login(a.TbAccounts, w, r)
}

func (a *App) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	Account.DeleteAccount(a.TbAccounts, w, r)
}

func (a *App) ListAccounts(w http.ResponseWriter, r *http.Request) {
	Account.ListAccounts(a.TbMovies, w, r)
}

func (a *App) ListMovies(w http.ResponseWriter, r *http.Request) {
	Movies.ListMovies(a.TbMovies, w, r)
}

func (a *App) Optional(w http.ResponseWriter, r *http.Request) {
	Optional.Optional(w, r)
}

func (a *App) Run(host string) {
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "Access-Control-Allow-Origin", "x-access-token"})
	corsObj := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(host, handlers.CORS(corsObj, headersOk, methodsOk)(a.Router)))
}
