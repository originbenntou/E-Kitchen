package template

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"html/template"
	"net/http"
	"time"

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
		"trClass": func(a pbShop.Status) string {
			if a == pbShop.Status_PRIVATE {
				return "uk-tile-muted"
			}
			if a == pbShop.Status_DELETED {
				return "uk-hidden"
			}
			return ""
		},
		"convertTime": func(a *timestamp.Timestamp) string {
			t, _ := ptypes.Timestamp(a)
			jst := time.FixedZone("Asia/Tokyo", 9*60*60)
			return t.In(jst).Format("2006/01/02 15:04")
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
