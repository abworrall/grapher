package gamegraph

import(
	"fmt"
)

type GameGraph struct {
	M map[Player][]int
}

type Player string

func New() *GameGraph {
	gg := GameGraph{
		M: map[Player][]int{},
	}
	return &gg
}

func (gg *GameGraph)String() string {
	str := "GG{\n"
	for k, v := range gg.M {
		str += fmt.Sprintf(" %-8.8s: %v\n", string(k), v)
	}
	return str + "}\n"
}

func (gg *GameGraph)Add(p string, vals ...int) {
	pl := Player(p)
	if _,exists := gg.M[pl]; !exists {
		gg.M[pl] = []int{}
	}

	if len(gg.M[pl]) == 0 && vals[0] != 1 {
		gg.M[pl] = append(gg.M[pl], 1)
	}
	
	for _,val := range vals {
		gg.M[pl] = append(gg.M[pl], val)
	}
}
