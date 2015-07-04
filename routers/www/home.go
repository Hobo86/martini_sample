package www

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessionauth"
	"github.com/martini-contrib/sessions"

	"../../models"
)

func HomeHandler(user sessionauth.User, res http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "www/home", "")
}

func LoginHandler(session sessions.Session, user sessionauth.User, res http.ResponseWriter, req *http.Request, r render.Render) {
	if user.IsAuthenticated() {
		r.Redirect("/")
		return
	}
	url := req.Header.Get("REFERER")
	r.HTML(200, "www/login", url)
}

func LoginPostHandler(session sessions.Session, postedUser model.User, r render.Render, req *http.Request, db gorm.DB) {

	// session.Set("userName", "Steven")
	// session.Get("userName")

	user := &model.User{}
	u := user.GetUserByNicknamePwd(postedUser.Nickname, postedUser.Password, db)

	if u != nil {
		err := sessionauth.AuthenticateSession(session, u)
		if err != nil {
			r.JSON(500, err)
		}

		qParams := req.URL.Query()
		redirect := qParams.Get(sessionauth.RedirectParam)
		r.Redirect(redirect)
		return
	} else {
		r.Redirect(sessionauth.RedirectUrl)
		return
	}
}

func LogoutHandler(session sessions.Session, user sessionauth.User, r render.Render) {
	sessionauth.Logout(session, user)
	r.Redirect("/")
}

func RegistHandler(user sessionauth.User, res http.ResponseWriter, req *http.Request, r render.Render) {
	if user.IsAuthenticated() {
		r.Redirect("/")
		return
	}

	url := req.Header.Get("REFERER")
	r.HTML(200, "www/regist", url)
}

func RegistPostHandler(session sessions.Session, postedUser model.User, r render.Render, req *http.Request, db gorm.DB) {
	qParams := req.URL.Query()
	redirect := qParams.Get(sessionauth.RedirectParam)
	r.Redirect(redirect)
}
