package main

type Book struct {
	title   string
	content *Trie
}

func (b Book) Contains(word string) bool {
	return b.content.Contains(word)
}
