package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", HomeHandler)
	router.GET("/view/:domain/:url", ViewHandler)
	router.GET("/view/:domain/:url/html", HtmlHandler)
	router.GET("/view/:domain/:url/plain", PlainHandler)
	router.POST("/resend/:domain", ResendHandler)
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	log.Fatal(http.ListenAndServe(":4000", handlers.CombinedLoggingHandler(os.Stdout, router)))
}
