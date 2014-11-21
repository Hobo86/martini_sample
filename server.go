package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"

	"./app"
	"./app/admin"
	"./app/api"
)

func main() {
	m := martini.Classic()

	// 静态资源
	m.Use(martini.Static("assets"))

	// m.Use(Auth)

	// martini-contrib/auth
	m.Use(auth.Basic("steven", "123456"))

	// martini-contrib/render
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", app.HomeHandler)

	m.Get("/hello/:name", func(params martini.Params, r render.Render) {
		r.HTML(200, "hello", params["name"])
	})

	m.Get("/admin/user", admin.UserHandler)

	// 该方法将会在authorize方法没有输出结果的时候执行.
	// 用作接口及权限验证
	m.Get("/api/user", authorize, api.UserHandler)

	m.Get("/api/user/login", api.UserLoginHandler)

	m.Get("/api/user/regist", api.UserRegistHandler)

	m.Run()
}

func Auth(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("API-KEY") != "secret123" {
		res.WriteHeader(http.StatusUnauthorized)
		http.Error(res, "Nope", 401)
	}
}

func authorize(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("API-KEY") != "secret123" {
		res.WriteHeader(http.StatusUnauthorized)
		http.Error(res, "Nope", 401)
	} else {
		return
	}
}
