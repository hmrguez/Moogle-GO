package main

type Book struct {
	title   string
	content *Trie
}

func (b Book) Contains(word string) bool {
	return b.content.Contains(word)
}

func Contains(books []Book, word string) []string {
	var answer []string

	for _, book := range books {
		if book.Contains(word) {
			answer = append(answer, book.title)
		}
	}

	return answer
}
