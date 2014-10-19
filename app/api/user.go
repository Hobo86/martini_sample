package api

import (
	"net/http"

	"github.com/martini-contrib/render"
)

func UserHandler(res http.ResponseWriter, req *http.Request, r render.Render) {

	r.JSON(200, map[string]interface{}{"URI": "api user"})
}

func UserLoginHandler(res http.ResponseWriter, req *http.Request, r render.Render) {

	r.JSON(200, map[string]interface{}{"URI": "api user login"})
}

func UserRegistHandler(res http.ResponseWriter, req *http.Request, r render.Render) {

	r.JSON(200, map[string]interface{}{"URI": "api user regist"})
}
