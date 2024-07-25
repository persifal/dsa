package main

import "fmt"

func main() {
	t := newTrie()
	t.Insert("apple")
	fmt.Printf("%v\n", t.Search("apple"))
}

type Trie struct {
	root *node
}

type node struct {
	t    map[rune]*node
	term bool
}

func newNode() *node {
	return &node{
		make(map[rune]*node),
		false,
	}
}

func newTrie() Trie {
	return Trie{
		newNode(),
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for _, r := range word {
		if _, ok := curr.t[r]; !ok {
			curr.t[r] = newNode()
		}

		curr = curr.t[r]
	}

	curr.term = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, r := range word {
		if _, ok := curr.t[r]; ok {
			curr = curr.t[r]
		} else {
			return false
		}
	}

	return curr.term
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for _, r := range prefix {
		if _, ok := curr.t[r]; ok {
			curr = curr.t[r]
		} else {
			return false
		}
	}

	return true
}
