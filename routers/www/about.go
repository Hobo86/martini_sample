package www

import (
	"net/http"

	"github.com/martini-contrib/render"
)

func AboutHandler(res http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "www/about", "")
}
