package website

import (
    "net/http"

    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/acceptlang"
    "github.com/codegangsta/martini-contrib/render"
)

func init() {
    m := martini.Classic()
    m.Use(acceptlang.Languages())
    m.Use(render.Renderer())

    m.Get("", func (r render.Render, languages acceptlang.AcceptLanguages) {
        for _, language := range languages {
            switch language.Language {
            case "en-US", "en":

                langs := make(map[string]string)
                langs["en"]=""
                langs["pt"]="hidden"

                r.HTML(200, "index", langs)
                return
            case "pt-BR", "pt":
                langs := make(map[string]string)
                langs["en"]="hidden"
                langs["pt"]=""
                r.HTML(200, "index", langs)
                return
            }
        }

        // default language
        langs := make(map[string]string)
        langs["en"]=""
        langs["pt"]="hidden"
        r.HTML(200, "index", langs)
    })

    http.Handle("/", m)
}
