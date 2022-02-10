package main

import(
	"context"
	"fmt"
	"net/http"
	"strconv"

	hw "github.com/skypies/util/handlerware"

	"github.com/abworrall/grapher/pkg/gamegraph"
)

var(
	players = map[string]string{
		"A": "[A]d",
		"D": "[D]imey",
		"J": "M[J]Lewis",
		"M": "[M]att",
		"L": "[L]aura",
		"P": "[P]ete",
	}
)

// /?p=person    id of the player
//  &move=123    next move to add
//
func formHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	templates := hw.GetTemplates(ctx)
	var params = map[string]interface{}{
		"Action": "/",
		"Players": players,
	}

	p := r.FormValue("p")
	if p == "" {
		templates.ExecuteTemplate(w, "form-name", params)
		return
	}

	params["Player"] = p
	params["PlayerName"] = players[p]
	params["VizAction"] = "/graph?p=" + p
	params["UndoAction"] = "/undo?p=" + p

	if mv := r.FormValue("move"); mv != "" {
		m,_ := strconv.Atoi(mv)
		gg := ReadGameGraph(ctx)
		gg.Add(p, m)
		WriteGameGraph(ctx, gg)
	}
	
	if err := templates.ExecuteTemplate(w, "form", params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func vizHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	templates := hw.GetTemplates(ctx)
	var params = map[string]interface{}{
		"VizAction": "/graph",
	}
	if err := templates.ExecuteTemplate(w, "viz", params); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func graphHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "inline; filename=\"WARLOCK.png\"")

	gg := ReadGameGraph(ctx)

	if r.FormValue("p") != "" {
		gg.RenderPlayer(w, r.FormValue("p"))
	} else {
		gg.RenderAllPlayers(w)
	}
}

func resetHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	gg := ReadGameGraph(ctx)

	if p := r.FormValue("p"); p != "" {
		delete(gg.M, gamegraph.Player(p))
	} else {
		gg = gamegraph.New()
	}

	WriteGameGraph(ctx, gg)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("reset complete")))
}

func undoHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	p := r.FormValue("p");
	if p == "" {
		http.Error(w, "Need a `&p=A` player for /undo", http.StatusInternalServerError)
		return
	}
	pl := gamegraph.Player(p)
	
	gg := ReadGameGraph(ctx)
	gg.M[pl] = gg.M[pl][:len(gg.M[pl])-1] // slice pop
	WriteGameGraph(ctx, gg)

	http.Redirect(w, r, fmt.Sprintf("/?p=%s", p), http.StatusFound)
}
