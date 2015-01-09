package website

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

//func init() {
//m := martini.Classic()
//m.Use(acceptlang.Languages())
//m.Use(render.Renderer())

//m.Get("/:lang", func(params martini.Params, r render.Render, languages acceptlang.AcceptLanguages) {
//lang := params["lang"]

//langSupported := map[string]bool{
//"pt": true,
//"en": true,
//}

////if !langSupported[lang] {
////// try to get it from header
////for _, language := range languages {
////switch language.Language {
////case "en-US", "en":
////lang = "en"
////break
////case "pt-BR", "pt":
////lang = "pt"
////break
////}
////}
////}

//data := map[string]interface{}{}
//switch lang {
//case "en":
//r.HTML(200, "page_en", data)
//default:
//r.HTML(200, "page_pt", data)
//}
//})

//http.Handle("/", m)
//}

var (
	DefaultPort = "8080"

	englishTemplate    = template.Must(template.ParseFiles("templates/page_en.tmpl"))
	portugueseTemplate = template.Must(template.ParseFiles("templates/page_pt.tmpl"))
)

func homepage(res http.ResponseWriter, req *http.Request) {
	isEnglish := strings.HasPrefix(req.URL.Path, "/en")
	if isEnglish {
		err := englishTemplate.Execute(res, make(map[string]interface{}))
		if err != nil {
			log.Println("englishTemplate:", err)
		}
		return
	}

	err := portugueseTemplate.Execute(res, make(map[string]interface{}))
	if err != nil {
		log.Println("portugueseTemplate:", err)
	}

	log.Println("end")
}

func init() {
	// Static assets (CSS, JS, images)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))

	http.HandleFunc("/", homepage)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Printf("PORT environmental variable is not set. Using default port %v.\n", DefaultPort)
		port = DefaultPort
	}

	log.Printf("Listening at port %v...\n", port)
	err := http.ListenAndServe(":"+port, http.DefaultServeMux)
	if err != nil {
		log.Fatalln("Fatal: http.ListenAndServe:", err)
	}
}
