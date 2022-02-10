package main

import(
	"fmt"
	"net/http"
)

func debugHandler(w http.ResponseWriter, r *http.Request) {
	str := "Hello me\n"

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("OK!\ndicebot debug handler\nin  [%s]\nout [%s]\n", r.FormValue("q"), str)))
}
