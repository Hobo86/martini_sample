package main

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"

	"./router"
	"./router/admin"
	"./router/api"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db            gorm.DB
	sqlConnection string
)

func main() {
	var err error
	sqlConnection = "root:meet77102@tcp(127.0.0.1:3306)/im?parseTime=True"
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err)
		fmt.Print("test")
		return
	}

	m := martini.Classic()

	// 静态资源
	m.Use(martini.Static("assets"))

	m.Map(db)

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

	m.Get("/admin/user", admin.UserHandler)

	// 该方法将会在authorize方法没有输出结果的时候执行.
	// 用作接口及权限验证
	m.Group("/api", func(r martini.Router) {
		m.Get("/user", authorize, api.UserHandler)

		m.Get("/user/login", api.UserLoginHandler)

		m.Get("/user/regist", api.UserRegistHandler)
	})

	m.NotFound(func() {
		// 处理 404
	})

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
