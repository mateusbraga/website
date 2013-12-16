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

    m.Get("/js/script.js", func (r render.Render, languages acceptlang.AcceptLanguages) {
        for _, language := range languages {
            switch language.Language {
            case "en-US", "en":
                r.HTML(200, "script", "en-US")
                return
            case "pt-BR", "pt":
                r.HTML(200, "script", "pt-BR")
                return
            }
        }
        r.HTML(200, "script", "pt-BR")
    })

    http.Handle("/", m)
}
