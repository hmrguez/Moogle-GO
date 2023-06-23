package main

import (
	"strings"
)

func preprocessTexts(text string) *Trie {
	var splitFunction = func(text string) []string {
		return strings.FieldsFunc(text, func(r rune) bool {
			return r == '.' || r == ' ' || r == ','
		})
	}

	var splitText = splitFunction(text)
	return prepareTrie(splitText)
}

func prepareTrie(text []string) *Trie {
	var trie = NewTrie()

	for _, word := range text {
		trie.Insert(word)
	}

	return trie
}

func getBooks() []Book {
	var files = readFiles()
	var books []Book

	for _, file := range files {
		var trie = preprocessTexts(file[1])
		var newBook = Book{title: file[0], content: trie}
		books = append(books, newBook)
	}

	return books
}
