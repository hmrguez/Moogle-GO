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

func prepareTrie(text []string) *Trie{
	var trie = NewTrie()

	for _, word := range text{
		trie.Insert(word)
	}

	return trie
}