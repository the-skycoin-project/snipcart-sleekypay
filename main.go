// /* */ //
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/snipcart/sleekypay/statik"
	"github.com/rakyll/statik/fs"
  //"encoding/json"
//flags "github.com/spf13/pflag"
)
const port = 8041

//package gorilla	// Routing based on the gorilla/mux router
var Serve http.Handler

	func main(){
		statikFS, err := fs.New()
		if err != nil {
			log.Fatal(err)
		}
		r := mux.NewRouter() //.StrictSlash(true)
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(statikFS)))
		Serve = r
		fmt.Printf("listening on http://127.0.0.1:%d using gorilla router\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), Serve))
	}
