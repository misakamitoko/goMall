package trie

import "strings"

type PathTrie struct {
	root *Node
}

type Node struct {
	children map[string]*Node
	isEnd    bool
	server   string
}

func CreateNewPathTrie() *PathTrie {
	return &PathTrie{
		root: &Node{
			children: make(map[string]*Node),
			isEnd:    false,
		},
	}
}

func (t *PathTrie) Insert(path, server string) {
	paths := strings.Split(path, "/")
	node := t.root
	for _, v := range paths {
		if _, ok := node.children[string(v)]; !ok {
			node.children[string(v)] = &Node{
				children: make(map[string]*Node),
				isEnd:    false,
				server:   "",
			}
		}
		node = node.children[string(v)]
	}
	node.isEnd = true
	node.server = server
}

func (t *PathTrie) Search(path string) (string, bool) {
	paths := strings.Split(path, "/")
	node := t.root
	for _, v := range paths {
		if _, ok := node.children[string(v)]; !ok {
			return "", false
		}
		node = node.children[string(v)]
	}
	return node.server, node.isEnd
}
