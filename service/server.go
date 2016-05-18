package service

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/cf-tools"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()
	webClient := fulfillmentWebClient{
		rootURL: "http://localhost:3001/skus",
	}

	appEnv, err := cfenv.Current()
	if err == nil {
		val, err := cftools.GetVCAPServiceProperty("backing-fulfill", "url", appEnv)
		if err == nil {
			webClient.rootURL = val
		} else {
			fmt.Printf("Failed to get URL property from bound service: %v\n", err)
		}
	}
	fmt.Printf("Using the following URL for fulfillment backing service: %s\n", webClient.rootURL)

	initRoutes(mx, formatter, webClient)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, webClient fulfillmentClient) {
	mx.HandleFunc("/catalog", getAllCatalogItemsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/catalog/{sku}", getCatalogItemDetailsHandler(formatter, webClient)).Methods("GET")
}
