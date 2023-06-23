package main

import (
	"strings"
	"sync"
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
	files := readFiles()
	booksChan := make(chan Book)
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file [2]string) {
			defer wg.Done()
			trie := preprocessTexts(file[1])
			newBook := Book{title: file[0], content: trie}
			booksChan <- newBook
		}(file)
	}

	go func() {
		wg.Wait()
		close(booksChan)
	}()

	var books []Book
	for book := range booksChan {
		books = append(books, book)
	}

	return books
}
