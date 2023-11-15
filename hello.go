package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.H1().Text("Hello world!")
}
