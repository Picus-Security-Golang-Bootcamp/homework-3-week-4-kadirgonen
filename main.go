package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Books struct {
	Books []Book `json:books`
}
type Book struct {
	BookID      int     `json:"book_id" gorm:"primaryKey"`
	Names       string  `json:"names"`
	PaperNumber int     `json:"paper_number"`
	StockNumber int     `json:"stock_number"`
	Cost        float64 `json:"cost"`
	StockCode   int     `json:"stock_code"`
	Isbn        int     `json:"ISBN"`
	AuthorID    int
	Author      Author `gorm:"foreignKey:AuthorID;references:ID"`
	Deleted     gorm.DeletedAt
	Flag        bool `json:"flag"`
}

type Author struct {
	ID   int32 `json:"id" gorm:"primary_key"`
	Name string
}

func main() {
	jsonFile, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	values, _ := ioutil.ReadAll(jsonFile)
	bookJson := Books{}
	json.Unmarshal(values, &bookJson)

	fmt.Println("aog")
	dsn := "host=localhost user=postgres password=2236386alper dbname=PatikaGoDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate()

	db.Migrator().CurrentDatabase()
	db.Migrator().CreateTable(&Author{})
	db.Migrator().CreateTable(&Book{})

	for _, book := range bookJson.Books {
		db.FirstOrCreate(&book)
	}
	if os.Args[1] == "listbook" {

		GetAllBooksWithoutAuthorInformation(*db)

	} else if os.Args[1] == "listauthor" {

		GetAllAuthorsWithoutBookInformation(*db)

	} else if os.Args[1] == "first" {

		GetFirstBookInformation(*db)

	} else if os.Args[1] == "last" {

		GetLastBookInformation(*db)
	} else if os.Args[1] == "authorbook" {

		GetAllAuthorsWithBookInformation(*db)
	} else if os.Args[1] == "bookauthor" {

		GetAllBooksWithAuthorInformation(*db)
	} else if os.Args[1] == "search" {

		var searchArg = os.Args[2:]
		SearchBookName(*db, strings.Join(searchArg, " "))

	} else if os.Args[1] == "get" {

		var getArg []string = os.Args[2:]
		fmt.Printf("book_id : %+v", getArg)

		i1, err := strconv.Atoi(os.Args[2])
		if err == nil {
			fmt.Println(i1)
		}
		GetBookID(*db, i1)

	} else if os.Args[1] == "buy" {

		var getArg []string = os.Args[2:]
		fmt.Printf("book_stock : %+v", getArg)

		i1, err := strconv.Atoi(os.Args[2])
		if err == nil {
			fmt.Println(i1)
		}
		i2, err := strconv.Atoi(os.Args[3])
		if err == nil {
			fmt.Println(i2)
		}
		GetBookStock(*db, i1, i2)
	} else if os.Args[1] == "delete" {

		var searchArg []string = os.Args[2:]
		fmt.Printf("Names : %+v", searchArg)

		i1, err := strconv.Atoi(os.Args[2])
		if err == nil {
			fmt.Println(i1)
		}

		DeleteBookID(*db, i1)
	} else {
		fmt.Println("Entered command is not valid!")
		usage()
	}
}
func DeleteBookID(db gorm.DB, ID int) {
	var book Book
	db.Delete(&book, ID)
}
func GetBookStock(db gorm.DB, id int, number int) {
	var book []Book
	db.Where(Book{BookID: id}).Find(&book)
	for _, v := range book {
		if v.StockNumber >= number {
			v.StockNumber = v.StockNumber - number
		} else {
			log.Fatal(" Stock Not Found")
		}
	}
	db.Save(&book)
}
func SearchBookName(db gorm.DB, names string) {
	var bookstore []Book
	// LIKE
	db.Where("names ILIKE ?", "%"+names+"%").Find(&bookstore)
	for i := 0; i < len(bookstore); i++ {
		fmt.Println("ALL-BOOKS ----------> " + bookstore[i].Names)
	}
}
func GetAllBooksWithoutAuthorInformation(db gorm.DB) {
	var bookstore []Book
	db.Find(&bookstore)
	for i := 0; i < len(bookstore); i++ {
		fmt.Println("ALL-BOOKS --------> " + bookstore[i].Names)
	}
}
func GetAllAuthorsWithoutBookInformation(db gorm.DB) {

	var bookstore []Author
	db.Find(&bookstore)
	for i := 0; i < len(bookstore); i++ {
		fmt.Println("ALL-AUTHORS --------> " + bookstore[i].Name)
	}
}
func GetFirstBookInformation(db gorm.DB) {

	var bookstore Book
	db.First(&bookstore, 1) // find product with integer primary key
	fmt.Println("ID --------> " + strconv.FormatUint(uint64(bookstore.BookID), 10))
	fmt.Println("NAME ----> " + bookstore.Names)
}
func GetLastBookInformation(db gorm.DB) {

	var bookstore Book
	db.Last(&bookstore)
	fmt.Println("ID --------> " + strconv.FormatUint(uint64(bookstore.BookID), 10))
	fmt.Println("Name ------> " + bookstore.Names)
}
func GetAllBooksWithAuthorInformation(db gorm.DB) {
	var books []Book

	db.Preload("Author").Find(&books)
	for i := 0; i < len(books); i++ {
		fmt.Println("ALL-BOOKS --------> ", books[i])
	}
}
func usage() {
	fmt.Println("list: list books")
	fmt.Println("search: search a book")
	fmt.Println("get: get a book")
	fmt.Println("delete: delete a book")
	fmt.Println("buy: make a stock control")

}
func GetAllAuthorsWithBookInformation(db gorm.DB) {
	var author []Author
	db.Preload("Books").Find(&author)
	for i := 0; i < len(author); i++ {
		fmt.Println("ALL-BOOKS --------> ", author[i])
	}
}
func GetBookID(db gorm.DB, id int) {
	var bookstore []Book
	db.Where(Book{BookID: id}).Find(&bookstore)
	if len(bookstore) == 0 {
		log.Fatal("Not Found ID")
	} else {
		for i := 0; i < len(bookstore); i++ {
			fmt.Println("ALL-BOOKS --------> " + strconv.Itoa(bookstore[i].BookID) + " " + bookstore[i].Names)
		}
	}
}
