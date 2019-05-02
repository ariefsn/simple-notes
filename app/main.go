package main

import (
	Ctrl "notes/app/controllers"
	"notes/app/core"

	"github.com/eaciit/toolkit"

	"git.eaciitapp.com/sebar/knot"
)

func main() {
	_ = core.Configuration()
	// address := core.App().GetString("host") + ":" + core.API().GetString("port")

	app := knot.NewApp()
	_ = app.Register(&Ctrl.NoteController{}, "api/note")
	_ = app.Register(&Ctrl.CategoryController{}, "api/category")
	_ = app.Register(&Ctrl.TagController{}, "api/tag")
	_ = app.Register(&Ctrl.SettingController{}, "api/setting")

	s := knot.NewServer()
	if core.IsDev() {
		s.ReverseProxy("/", core.FullAddress(core.Public().GetString("port")), "VirtualDirectory")
	} else {
		// Set Public route for static directory, this is for read dist/css & dist/js
		// Can set in vue.config.js @ publicPath => optional
		app.Static("public", core.ViewsPath(""))
		s.Route("/", func(ctx *knot.WebContext) {
			err := ctx.App().SetViewsPath(core.ViewsPath(""))
			if err != nil {
				ctx.WriteJSON(err.Error(), 400)
			}
			ctx.WriteTemplate(toolkit.M{
				"title": core.App().GetString("title"),
			}, core.ViewsPath("index.html"))
		})
	}

	s.RegisterApp(app, "")
	// s.Start(address)
	s.Start(":" + core.API().GetString("port"))
	s.Wait()
}
