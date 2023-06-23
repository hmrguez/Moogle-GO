package main

import "fmt"

func main() {

	var books = getBooks()

	for {
		fmt.Println("Write a word: ")
		var word string
		fmt.Scan(&word)

		fmt.Println()
		fmt.Println("The books that contain the word are: ")

		var result = Contains(books, word)

		for _, res := range result {
			fmt.Println(res)
		}

		fmt.Println()
		fmt.Println("--------------------------------------------")
		fmt.Println()
	}
}
