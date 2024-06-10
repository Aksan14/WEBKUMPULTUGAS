package controller

import (
	"net/http"
	"strconv"
	"tesgolang/entities"
	"tesgolang/models"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categ := models.See()

	data := map[string]any{
		"category": categ,
	}

	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)

}

func Create(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/created.html")


		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var catego entities.User

		catego.NAMA = r.FormValue("name")
		catego.LINKTUGAS = r.FormValue("nama")
		catego.WAKTUKUMPUL = time.Now()

		if ok := models.Create(catego); !ok {
			temp, err := template.ParseFiles("views/created.html")

			if err != nil {
				panic(err)
			}

			temp.Execute(w, nil)
		}
		http.Redirect(w, r ,"/", http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request)  {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	
	if err != nil {
		panic(err)
	}

	if err := models.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}