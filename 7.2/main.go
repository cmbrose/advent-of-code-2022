package main

import (
	"fmt"
	"main/util"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	children map[string]*node
	parent   *node
}

func newNode(name string, size int, parent *node) *node {
	return &node{
		name:     name,
		size:     size,
		parent:   parent,
		children: make(map[string]*node),
	}
}

func printTree(n *node, indentation string) {
	fmt.Printf("%s- %s (%d)\n", indentation, n.name, n.size)

	for _, c := range n.children {
		printTree(c, indentation+"  ")
	}
}

func main() {
	root := newNode("/", 0, nil)

	currentNode := root

	for _, line := range util.ReadInputLines("./input.txt") {
		if strings.HasPrefix(line, "$ cd ") {
			target := strings.TrimPrefix(line, "$ cd ")

			if target == "/" {
				currentNode = root
			} else if target == ".." {
				currentNode = currentNode.parent
			} else if node, ok := currentNode.children[target]; ok {
				currentNode = node
			} else {
				node := newNode(target, 0, currentNode)
				currentNode.children[target] = node
				currentNode = node
			}
		} else if !strings.HasPrefix(line, "$ ls") {
			pair := strings.Split(line, " ")

			if pair[0] == "dir" {
				continue
			}

			size, err := strconv.Atoi(pair[0])
			util.Check(err)

			name := pair[1]

			if _, ok := currentNode.children[name]; ok {
				continue
			}

			node := newNode(name, size, currentNode)
			currentNode.children[name] = node

			// Propagate the size back up the tree
			for n := currentNode; n != nil; n = n.parent {
				n.size += size
			}
		}
	}

	printTree(root, "")

	threshold := root.size - 40000000

	smallestNode := root

	toWalk := []*node{root}

	for len(toWalk) > 0 {
		n := toWalk[0]
		toWalk = toWalk[1:]

		if len(n.children) == 0 {
			continue // only consider dirs
		}

		if n.size > threshold && n.size < smallestNode.size {
			smallestNode = n
		}

		for _, c := range n.children {
			toWalk = append(toWalk, c)
		}
	}

	fmt.Printf("%d (from %s)\n", smallestNode.size, smallestNode.name)
}
