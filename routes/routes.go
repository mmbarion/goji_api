package routes

import (
	"../apps/CollectingCheckList/routes"
	"../controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Include() {
	goji.Get("/",controllers.Home)
	goji.Get("/about/:name",controllers.About)

	//CollectingCheckList app
	collectingchecklist := web.New()
	routes.CollectingCheckList(collectingchecklist)
}