package main

import (
	"fmt"
)

type Tree struct {
	Next map[rune]*Tree
	Data rune
	Size int
}

func TreeNew(r rune) *Tree {
	t := &Tree{Next: make(map[rune]*Tree), Data: r, Size: 1}
	return t
}

func (t *Tree) Add(s string) {
	parent, ok := t.Next[rune(s[0])]
	if ok {
		parent.Size += 1
	} else {
		parent = TreeNew(rune(s[0]))
		t.Next[rune(s[0])] = parent
	}
	for i := 1; i < len(s); i++ {
		curr, ok := parent.Next[rune(s[i])]
		if ok {
			curr.Size += 1
		} else {
			curr = TreeNew(rune(s[i]))
			parent.Next[rune(s[i])] = curr
		}
		parent = curr
	}
}

func Print(t *Tree) {
	fmt.Println(t)
	for _, tree := range t.Next {
		Print(tree)
	}
}

func (t *Tree) Find(s string) int {
	parent, ok := t.Next[rune(s[0])]
	if !ok {
		return 0
	}

	for i := 1; i < len(s); i++ {
		curr, ok := parent.Next[rune(s[i])]
		if !ok {
			return 0
		}
		if i == len(s)-1 {
			return curr.Size
		}
		parent = curr
	}
	return parent.Size
}

func main() {
	root := &Tree{Size: -1}
	root.Next = make(map[rune]*Tree)

	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var cmd, value string
		fmt.Scanf("%s", &cmd)
		fmt.Scanf("%s", &value)
		if cmd == "add" {
			root.Add(value)
		} else if cmd == "find" {
			fmt.Println(root.Find(value))
		}
	}
}
