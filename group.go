package server

import "github.com/dinozor-io/interfaces"

type Group struct {
	condition func(interfaces.Controller) bool
}

func (g *Group) Cond(cond func(interfaces.Controller) bool) {
	g.condition = cond
}

func (g *Group) CheckCond(c interfaces.Controller) bool {
	return g.condition(c)
}
