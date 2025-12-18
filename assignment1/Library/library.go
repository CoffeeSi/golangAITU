package Library

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	ID         uint64
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[uint64]Book
}

func (library *Library) AddBook(book Book) {
	if library.Books == nil {
		library.Books = make(map[uint64]Book)
	}
	library.Books[book.ID] = book
	fmt.Printf("Book %s successfully added!\n", book.Title)
}

func (library *Library) BorrowBook(book Book) {
	var currentBook = library.Books[book.ID]
	if !currentBook.IsBorrowed && currentBook == book {
		currentBook.IsBorrowed = true
		library.Books[book.ID] = currentBook
		fmt.Printf("Book %s successfully borrowed!\n", currentBook.Title)
	} else {
		fmt.Printf("Book %s already borrowed or not exists!\n", book.Title)
	}
}

func (library *Library) ReturnBook(book Book) {
	var currentBook = library.Books[book.ID]
	if currentBook.IsBorrowed && currentBook == book {
		currentBook.IsBorrowed = false
		library.Books[book.ID] = currentBook
		fmt.Printf("Book %s successfully returned!\n", currentBook.Title)
	} else {
		fmt.Printf("Book %s already returned or not match!\n", book.Title)
	}
}

func (library *Library) ListAvailableBooks() {
	if len(library.Books) == 0 {
		fmt.Println("The library is empty!")
		return
	}
	for i, book := range library.Books {
		fmt.Printf("%v. ID: %v, Title: %s, Author: %s, IsBorrowed: %t\n", i+1, book.ID, book.Title, book.Author, book.IsBorrowed)
	}
}

func (library *Library) ConsoleMenu() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var id uint64 = 0
	for {
		fmt.Print("\nInput operation:\n 1. Add book\n 2. Borrow book\n 3. Return book \n 4. List available books \n 0. Exit\n: ")
		if !scanner.Scan() {
			return
		}
		operation := scanner.Text()
		switch operation {
		case "1":
			var title, author string
			fmt.Print("Book name: ")
			scanner.Scan()
			title = scanner.Text()
			fmt.Print("Book author: ")
			scanner.Scan()
			author = scanner.Text()
			newBook := Book{
				ID:         id,
				Title:      title,
				Author:     author,
				IsBorrowed: false,
			}
			id++
			library.AddBook(newBook)
		case "2":
			var bookId uint64
			fmt.Print("ID of book to borrow: ")
			scanner.Scan()
			bookId, _ = strconv.ParseUint(scanner.Text(), 10, 64)
			book, exists := library.Books[bookId]
			if !exists {
				fmt.Println("Book does not exist")
				continue
			}
			library.BorrowBook(book)
		case "3":
			var bookId uint64
			fmt.Print("ID of book to return: ")
			scanner.Scan()
			bookId, _ = strconv.ParseUint(scanner.Text(), 10, 64)
			book := library.Books[bookId]
			library.ReturnBook(book)
		case "4":
			library.ListAvailableBooks()
		case "0":
			return
		default:
			fmt.Println("Invalid number of operation!")
		}
	}
}
