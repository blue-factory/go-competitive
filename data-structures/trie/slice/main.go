package main

import "fmt"

const (
	alphabetLetters = 26
	firstLetter     = 'a'
)

type trieNode struct {
	children [alphabetLetters]*trieNode
	isWord   bool
}

// Trie ...
type Trie struct {
	root *trieNode
}

// Init ...
func Init() *Trie {
	return &Trie{
		root: &trieNode{},
	}
}

func (t *Trie) insert(word string) {
	current := t.root

	for i, w := range word {
		index := w - firstLetter

		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}

		current = current.children[index]
	}

	current.isWord = true
}

func (t *Trie) find(word string) bool {
	current := t.root

	for _, w := range word {
		index := w - firstLetter

		if current.children[index] == nil {
			return false
		}

		current = current.children[index]
	}

	if !current.isWord {
		return true
	}

	return false
}

func main() {
	trie := Init()
	words := []string{"abc", "ledo", "dog"}

	for _, w := range words {
		trie.insert(w)
		fmt.Println(fmt.Sprintf("insert word=%s", w))
	}

	wordsToFind := []string{"ab", "led", "abc", "ledo", "dog", "dogs", "og"}

	for _, w := range wordsToFind {
		ok := trie.find(w)
		if ok {
			fmt.Println(fmt.Sprintf("found in tree word=%s", w))
			continue
		}

		fmt.Println(fmt.Sprintf("not found in tree word=%s", w))
	}
}
