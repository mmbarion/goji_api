package routes

import (
	"../controllers"
	"net/http"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func CollectingCheckList(m *web.Mux) {
	goji.Handle("/collectingchecklist/*",m)
	goji.Get("/collectingchecklist",http.RedirectHandler("/collectingchecklist/",301))

	m.Get("/collectingchecklist/",controllers.CCLHome)
}