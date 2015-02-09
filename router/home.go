package router

import (
	"net/http"

	"../model"
	"github.com/martini-contrib/render"
)

func HomeHandler(res http.ResponseWriter, req *http.Request, r render.Render) {
	db.First(&user)
	// r.HTML(200, "home", "Steven")
	r.JSON(200, user)
}
