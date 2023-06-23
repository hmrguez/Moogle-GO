package main

import "sync"

type Book struct {
	title   string
	content *Trie
}

func (b Book) Contains(word string) bool {
	return b.content.Contains(word)
}

func Contains(books []Book, word string) []string {
	answerChan := make(chan string)
	var wg sync.WaitGroup

	for _, book := range books {
		wg.Add(1)
		go func(book Book) {
			defer wg.Done()
			if book.Contains(word) {
				answerChan <- book.title
			}
		}(book)
	}

	go func() {
		wg.Wait()
		close(answerChan)
	}()

	var answer []string
	for bookTitle := range answerChan {
		answer = append(answer, bookTitle)
	}

	return answer
}
