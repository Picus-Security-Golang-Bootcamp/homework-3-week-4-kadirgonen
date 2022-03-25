# Homework 3 Week 4

In this application, we provide a few status checks of the books in the application with the help of database by making a book application.

## Purpose of application

We have a list of books.
The book areas are as follows;
```
- Book ID
- Book Name
- Paper Number
- Stock Number
- Cost
- Stock Code
- ISBN
- Author (ID and Name)
```
1. Book and Author information read from a file.
2. It record in the database.
3. Database queries writte.(GORM Queries)
4. GORM Queries writte(
5. List all books and all authors (listbook and listauthor)
6. It brings the first book.
7. It brings the last book.
8. List the books in which the given entry is in the titles of the books (search)
9. Delete the book with the given ID. (The deleted book should be coming by ID.) (delete)
10. Print book by ID(get)
11. Buy the book with the ID given as many as you want and print the last information of the book on the screen. (buy)
12. If the wrong command is entered, it will print the usage on the screen. 
 
## Used Commands

### listbook command
```
go run main.go listbook
```
This command allows us to see the books we define in the application as a list.

### listauthor command
```
go run main.go listbook
```
This command allows us to see the authors we define in the application as a list.

### first command
```
go run main.go first
```
This command allows us to see the first book in the list of the books we have defined in the application.

### last command
```
go run main.go last
```
This command allows us to see the last book in the list of the books we have defined in the application.

### search command 
```
go run main.go search <bookName>
go run main.go search Harry Potter
```
This command allows us to scan the objects we have defined in the application to check if they exist inside the application.

### get command
```
go run main.go get <bookID>
go run main.go get 6
```
This command allows to bring us the objects used in the application and their contents.

### buy command
```
go run main.go buy <bookID> <quantity>
go run main.go buy 3 4
```
This command checks the stock count of objects used in the application. It checks the stock of the items we will purchase later and shows us the number of stocks after the purchase accordingly.

### delete command
```
go run main.go delete <bookID>
go run main.go delete 4
```
This command allows us to delete objects used in the application.

## Requirements

* Go Language
* Git
* Go Module
* GORM

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
