package main

import (
   "github.com/urfave/negroni"
   "github.com/aloksinghal/Sentrymiddleware"
    "github.com/gorilla/mux"
    raven "github.com/getsentry/raven-go"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    panic("help")
}

func main() {
	raven.SetDSN("add your Sentry DSN here")
	n := negroni.New()
	n.Use(negroni.NewLogger())
	router := mux.NewRouter()
    n.Use(sentrymiddleware.Middleware{})
    n.UseHandler(router)
    router.HandleFunc("/test", handler)
    http.ListenAndServe("0.0.0.0:8080", n)

}