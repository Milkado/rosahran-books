package controllers

import (
	"github/milkado/rosharan-books/model"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	books := model.RetrieveAllBooks()
	temp.ExecuteTemplate(w, "Index", books)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		author := r.FormValue("author")
		blurb := r.FormValue("blurb")
		price, errPrice := strconv.ParseFloat(r.FormValue("price"), 64)
		stock, errStock := strconv.Atoi(r.FormValue("stock"))

		if errPrice != nil {
			panic(errPrice.Error())
		}

		if errStock != nil {
			panic(errStock.Error())
		}

		model.CreateBook(title, author, blurb, price, stock)
	}
	http.Redirect(w, r, "/", 301)
}

func Edit (w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	book := model.ShowBook(id)
	temp.ExecuteTemplate(w, "Edit", book)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, errId := strconv.Atoi(r.FormValue("id"))
		title := r.FormValue("title")
		author := r.FormValue("author")
		blurb := r.FormValue("blurb")
		price, errPrice := strconv.ParseFloat(r.FormValue("price"), 64)
		stock, errStock := strconv.Atoi(r.FormValue("stock"))

		if errId != nil {
			panic(errId.Error())
		}

		if errPrice != nil {
			panic(errPrice.Error())
		}

		if errStock != nil {
			panic(errStock.Error())
		}

		model.UpdateBook(id, title, author, blurb, price, stock)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	model.DeleteBook(id)

	http.Redirect(w, r, "/", 301)
}
