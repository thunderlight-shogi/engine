package website

import website "github.com/thunderlight-shogi/engine/website/controllers"

var Routes = []route{
	Get("/", website.Home),
	Get("/editor/{id}", website.Editor),
}
