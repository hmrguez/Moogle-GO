package main

import "fmt"

// "fmt"

func main() {
    var files = readFiles()
	var books []Book

	for _, file := range files{
		var trie = preprocessTexts(file[1])
		var newBook = Book{title: file[0], content: trie}
		books = append(books, newBook);
	}

	fmt.Println(books[0].Contains("Shakespeare"))
	fmt.Println(books[0].Contains("Turing"))
	fmt.Println(books[0].Contains("Romeo"))
	fmt.Println(books[0].Contains("Juliet"))
}