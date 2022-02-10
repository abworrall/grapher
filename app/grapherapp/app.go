package main

// UI: withctx handler


// 3. cloud datastore singleton for GameGraph
// 4. UI for adding players/nodes, and serializing

import(
	"fmt"
	"log"
	"net/http"
	"os"

	hw "github.com/skypies/util/handlerware"
)

func init() {
	hw.RequireTls = false
	hw.InitTemplates("app/web/templates") // relative to go module root, which is git repo root

	http.HandleFunc("/",             hw.WithCtx(formHandler))
	http.HandleFunc("/viz",          hw.WithCtx(vizHandler))
	http.HandleFunc("/reset",        hw.WithCtx(resetHandler))
	http.HandleFunc("/undo",         hw.WithCtx(undoHandler))
	http.HandleFunc("/graph",        hw.WithCtx(graphHandler))

	log.Printf("(init has run)\n")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("[grapherapp] listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
