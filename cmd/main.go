package main

import (
	"Library_CRM/storage/jsons"
	"Library_CRM/storage/repo"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


func main() {
	var Books jsons.BookLibrary
	data, err := ioutil.ReadFile("../storage/jsons/books.json")
	if err != nil {
		fmt.Println("Error reading:",err)
	}
	err = json.Unmarshal(data,&Books)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	var response  int

	for response != 7 {
		fmt.Println("---------------------------------------------------------------------------------------------")
		fmt.Println("-------------------Good to see you,Admin. what do you want to do?----------------------------")
		fmt.Println("---------------------------------------------------------------------------------------------")
		fmt.Println("<1> - See all the books' information")
		fmt.Println("<2> - Add the book to the library")
		fmt.Println("<3> - Remove the book from library")
		fmt.Println("<4> - Update the book's information")
		fmt.Println("<5> - Get the book by its ID")
		fmt.Println("<6> - Get the book by its Category")
		fmt.Println("<7> - Exit")
		fmt.Println("---------------------------------------------------------------------------------------------")

		fmt.Println("Enter your choice:")
		fmt.Scanln(&response)

		var books  []repo.Book
		
		reader := bufio.NewReader(os.Stdin)
		switch response{
		case 1:
			books = Books.GetBooks()
			for _, book := range books {
fmt.Println("---------------------------------------------------------------------------------------------------------")
fmt.Printf(`
Id : %v
title: %v
author: %v
year: %v
status: %v
price: %v
period: %v
category: %v
page: %v
`,book.ID,book.Title,book.Author,book.Year,book.Status,book.Price,book.Period,book.Category,book.Page)
}
		case 2:
			var book repo.Book
			book.ID = len(Books.Books) + 1
			
			
			
			fmt.Println("Enter the book's Title:")
			Title, _ := reader.ReadString('\n')
			book.Title = Title[:len(Title)-1]
			fmt.Println("Enter the book's Author:")
			Author, _ := reader.ReadString('\n')
			book.Author = Author[:len(Author)-1]
			fmt.Println("Enter the book's year:")
			fmt.Scanln(&book.Year)
			book.Status = "Given"
			fmt.Println("Enter the book's price:")
			fmt.Scanln(&book.Price)
			fmt.Println("Enter the book's period given:")
			fmt.Scanln(&book.Period)
			fmt.Println("Enter the book's category:")
			Category, _ := reader.ReadString('\n')
			book.Category = Category[:len(Category)-1]
			fmt.Println("Enter the book's pages:")
			fmt.Scanln(&book.Page)

			Books.Books = append(Books.Books, book)
		case 3:
			var id int
			fmt.Println("Enter the book's ID:")
			fmt.Scanln(&id)
			for i, book := range Books.Books {
				if id == book.ID {
					Books.Books = append(Books.Books[:i],Books.Books[i+1:]...)
				}
			}
		case 4:
			var id int
			fmt.Println("Enter the book's ID:")
			fmt.Scanln(&id)
			for i, book := range Books.Books {
				if id == book.ID {
					fmt.Println("Enter the book's info(available or given):")
					Status, _ := reader.ReadString('\n')
					Books.Books[i].Status = Status[:len(Status)-1]
					fmt.Println("Enter the book's new price:")
					fmt.Scanln(&Books.Books[i].Price)
					fmt.Println("Enter the book's new period:")
					fmt.Scanln(&Books.Books[i].Period)

				}
			}
		case 5:
			var id int
			var book repo.Book
			fmt.Println("Enter the book's ID:")
			fmt.Scanln(&id)
			for i, b := range Books.Books {
				if id == b.ID {
					book = Books.Books[i]
				}
			}
fmt.Println("---------------------------------------------------------------------------------------------------------")
fmt.Printf(`
Id : %v
title: %v
author: %v
year: %v
status: %v
price: %v
period: %v
category: %v
page: %v
`,book.ID,book.Title,book.Author,book.Year,book.Status,book.Price,book.Period,book.Category,book.Page)
		case 6:
			fmt.Println("Enter the book's Category:")
			Category,_ := reader.ReadString('\n')
			for _, book := range Books.Books {
				if Category == book.Category {
fmt.Println("---------------------------------------------------------------------------------------------------------")
fmt.Printf(`
Id : %v
title: %v
author: %v
year: %v
status: %v
price: %v
period: %v
category: %v
page: %v
`,book.ID,book.Title,book.Author,book.Year,book.Status,book.Price,book.Period,book.Category,book.Page)
				}
			}
		case 7:
			fmt.Println("\nYou have got out of the program :)")
		}
	}
	data,_ = json.Marshal(Books)
	err = ioutil.WriteFile("../storage/jsons/books.json", data, 0644)

}
