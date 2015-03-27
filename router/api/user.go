package api

import (
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"

	"../../model"
)

func UserHandler(params martini.Params, res http.ResponseWriter, req *http.Request, r render.Render, db gorm.DB) {

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	user := &model.User{}
	u := user.GetUserById(id, db)
	r.JSON(200, u)
}

func UserLoginHandler(res http.ResponseWriter, req *http.Request, r render.Render) {

	r.JSON(200, map[string]interface{}{"URI": "api user login"})
}

func UserRegistHandler(res http.ResponseWriter, req *http.Request, r render.Render) {

	r.JSON(200, map[string]interface{}{"URI": "api user regist"})
}