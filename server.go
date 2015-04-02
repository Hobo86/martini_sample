package main

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"

	"./config"
	"./router"
	"./router/admin"
	"./router/api"
)

func main() {

	m := martini.Classic()

	// 静态资源
	m.Use(martini.Static("assets"))

	m.Map(config.DB())

	// m.Use(Auth)

	// martini-contrib/auth
	m.Use(auth.Basic("steven", "123456"))

	// martini-contrib/render
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", router.HomeHandler)

	m.Get("/hello/:name", func(params martini.Params, r render.Render) {
		r.HTML(200, "hello", params["name"])
	})

	m.Get("/admin/user/(?P<id>[0-9]*)", admin.UserHandler)

	// 该方法将会在authorize方法没有输出结果的时候执行.
	// 用作接口及权限验证
	m.Group("/api", func(r martini.Router) {

		// m.Get("/user/:id", authorize, api.UserHandler)
		m.Get("/user/(?P<id>[0-9]*)", api.UserHandler)

		m.Get("/user/login", api.UserLoginHandler)

		m.Get("/user/regist", api.UserRegistHandler)
	})

	m.NotFound(func(r render.Render) {
		// 404
		r.HTML(404, "40x", "404-Not found!")
	})

	m.Run()
}

func InitDB() {

}

func Auth(res http.ResponseWriter, req *http.Request, r render.Render) {
	if req.Header.Get("API-KEY") != "secret123" {
		res.WriteHeader(http.StatusUnauthorized)

		// 403
		r.HTML(403, "40x", "403-Access Forbidden!")
	}
}

func authorize(res http.ResponseWriter, req *http.Request, r render.Render) {
	if req.Header.Get("API-KEY") != "secret123" {
		res.WriteHeader(http.StatusUnauthorized)

		// 403
		r.HTML(403, "40x", "403-Access Forbidden!")
	} else {
		return
	}
}
