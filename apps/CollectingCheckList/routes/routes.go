package routes

import (
	"../controllers"
	//"net/http"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func CollectingCheckList(m *web.Mux) {
	goji.Handle("/api/*",m)
	//goji.Get("/collectingchecklist",http.RedirectHandler("/collectingchecklist/",301))
	m.Use(middleware.SubRouter)
	// /api/collectingchecklist/
	m.Get("/collectingchecklist",controllers.CCLHome)
	m.Post("/collectingchecklist/addchecklist",controllers.AddCheckList)
	m.Get("/collectingchecklist/getstaff",controllers.GetStaff)
}