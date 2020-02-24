package template

import (
	"html/template"
	"net/http"

	pbShop "github.com/originbenntou/E-Kitchen/proto/shop"
)

func Render(w http.ResponseWriter, name string, content interface{}) {
	// 独自関数
	funcMap := template.FuncMap{
		"checked": func(a pbShop.Status, b string) string {
			if a.String() == b {
				return "checked"
			}
			return ""
		},
	}

	t := template.Must(template.New("t").Funcs(funcMap).ParseFiles(
		"/template/layout.tpl",
		"/template/header.tpl",
		"/template/"+name+".tpl"))

	if err := t.ExecuteTemplate(w, "layout", content); err != nil {
		panic(err)
	}
}
