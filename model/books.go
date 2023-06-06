package model

import (
	"github/milkado/rosharan-books/db"
)

type Book struct {
	Id     int
	Title  string
	Author string
	Blurb  string
	Price  float64
	Stock  int
}

func RetrieveAllBooks() []Book {
	db := db.ConnectDataBase()

	selectAll, err := db.Query("SELECT * FROM books")
	if err != nil {
		panic(err.Error())
	}

	b := Book{}
	books := []Book{}

	for selectAll.Next() {
		var id, stock int
		var title, author, blurb string
		var price float64

		err = selectAll.Scan(&id, &title, &author, &blurb, &price, &stock)

		if err != nil {
			panic(err.Error())
		}

		b.Id = id
		b.Title = title
		b.Author = author
		b.Blurb = blurb
		b.Price = price
		b.Stock = stock

		books = append(books, b)

	}
	defer db.Close()
	return books
}

func ShowBook(id string) Book {
	db := db.ConnectDataBase()

	selectOne, err := db.Query("SELECT * FROM books WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	b := Book{}

	for selectOne.Next() {
		var id, stock int
		var title, author, blurb string
		var price float64

		err = selectOne.Scan(&id, &title, &author, &blurb, &price, &stock)

		if err != nil {
			panic(err.Error())
		}

		b.Id = id
		b.Title = title
		b.Author = author
		b.Blurb = blurb
		b.Price = price
		b.Stock = stock

	}
	defer db.Close()
	return b
}

func UpdateBook(id int, title, author, blurb string, price float64, stock int) {
	db := db.ConnectDataBase()

	update, err := db.Prepare("UPDATE books SET title=?, author=?, blurb=?, price=?, stock=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(title, author, blurb, price, stock, id)
	defer db.Close()
}

func CreateBook(title, author, blurb string, price float64, stock int) {
	db := db.ConnectDataBase()

	insert, err := db.Prepare("INSERT INTO books(title, author, blurb, price, stock) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	insert.Exec(title, author, blurb, price, stock)
	defer db.Close()
}

func DeleteBook(id string) {
	db := db.ConnectDataBase()

	delete, err := db.Prepare("DELETE FROM books WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()
}
