package controllers

import "github.com/robfig/revel"

type Display struct {
	*revel.Controller
}

func (c Display) Index() revel.Result {
	return c.Render()
}
