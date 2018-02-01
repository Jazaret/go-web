package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	shopController shop
)

//Startup - This is the startup method
func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.standLocatorTemplate = templates["stand_locator.html"]
	homeController.loginTemplate = templates["login.html"]
	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
