package gamegraph

import(
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/goccy/go-graphviz"
)

func (gg *GameGraph)RenderAllPlayers(w io.Writer) error {
	return gg.render(w, "")
}

func (gg *GameGraph)RenderPlayer(w io.Writer, p string) error {
	return gg.render(w, p)
}

// If `p` is empty, render all players; else render player `p`
func (gg *GameGraph)render(w io.Writer, p string) error {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil { return err }

	defer func() {
		if err := graph.Close(); err != nil {
			log.Printf("WTF: %v\n", err)
			return
		}
		g.Close()
	}()

	labels := map[[2]int][]string{}

	if p != "" {
		graph.SetRankDir("BT")
	}

	for player, vals := range gg.M {
		if p != "" && Player(p) != player {
			continue
		}

		for i:=1; i<len(vals); i++ {
			u, v := vals[i-1], vals[i]
			k := [2]int{u,v}

			n1, err := graph.CreateNode(fmt.Sprintf("%d", u))
			if err != nil { log.Fatal("WTF3a: %v\n", err) }
			n2, err := graph.CreateNode(fmt.Sprintf("%d", v))
			if err != nil { log.Fatal("WTF3b: %v\n", err) }
			edge, err := graph.CreateEdge("", n1, n2)
			if err != nil { log.Fatal("WTF4: %v\n", err) }

			// If we're doing all players, build up edge labels
			if p == "" {
				if _,exists := labels[k]; !exists {
					labels[k] = []string{}
				}
				labels[k] = append(labels[k], string(player))

				label := strings.Join(labels[k], "")
				edge.SetLabel(label)
				edge.SetPenWidth(2 * float64(len(label)))
			}

			if v == 1 {
				edge.SetColor("red") // Character died
			}
		}
	}

	return g.Render(graph, graphviz.PNG, w)
}
