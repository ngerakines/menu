package controllers

import "github.com/robfig/revel"

type Create struct {
	*revel.Controller
}

func (c Create) Index() revel.Result {
	return c.Render()
}
