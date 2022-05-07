package application

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// queyr params are not actual query aorams. just substitue for any route

type QueryParams struct {
	Order  []string
	Params map[string]string
}

type CustomContext struct {
	params QueryParams
}

type HandlerFunction func(w http.ResponseWriter, r *http.Request, context CustomContext)

type Handler interface {
	HandlerFunc(string, string, HandlerFunction)
	ServeHTTP(http.ResponseWriter, *http.Request)
}
type route struct {
	method          string
	handlerFunction HandlerFunction
	queryParams     []string
}
type Router struct {
	routes map[string]route
	app    *Application
}

func (app *Application) Routes() Handler {
	var router *Router = &Router{
		routes: map[string]route{},
		app:    app,
	}

	var a = 3
	var _ = string(a)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/get", app.simpleGetHandler)
	router.HandlerFunc(http.MethodGet, "/movies", app.getMovies)
	router.HandlerFunc(http.MethodGet, "/movies/movie/:id", app.getMovieById)

	return router
}

func (r *Router) HandlerFunc(method string, routeUrl string, handlerFunction HandlerFunction) {
	routeUrl, queryParams := r.parseUrl(routeUrl)

	log.Printf("%s", routeUrl)

	if _, ok := r.routes[routeUrl]; ok {
		panic("route already exists")
	}

	r.routes[routeUrl] = route{
		method,
		handlerFunction,
		queryParams,
	}
}

func (r *Router) parseUrl(routeUrl string) (string, []string) {
	parsedUrl := strings.Split(routeUrl, ":")
	parsedUrl2 := []string(routeUrl)

	return parsedUrl[0], parsedUrl[1:]
}

// pokrenuti server
// za svaku rutu definiranu gore zvati http..hanbdleFunc()
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", router.app.config.Port), nil)
	if err != nil {
		log.Println("eror")
	}

	routeUrl, params := router.parseUrl(fmt.Sprintf("%s", r.URL))

	log.Printf("%s", routeUrl)

	method := r.Method
	route, ok := router.routes[routeUrl]

	if !ok {
		log.Println("no route")
		return
	}
	if method != route.method {
		log.Println("no method")
		return
	}

	queryParamsKeysLength := len(route.queryParams)
	queryParams := QueryParams{
		Order:  route.queryParams,
		Params: make(map[string]string),
	}
	for valueIndex, paramValue := range params {
		if valueIndex > queryParamsKeysLength {
			break
		}

		queryParams.Params[queryParams.Order[valueIndex]] = paramValue
	}

	context := CustomContext{
		params: queryParams,
	}

	log.Printf("%v", context)

	route.handlerFunction(w, r, context)
}
