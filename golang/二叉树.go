package main

import (
    "fmt"
)

type Tree struct {
    left *Tree
    right *Tree
    value int
}

type Node struct {
    leftChild *Node
    rightChild *Node
    value int
}

func (tree Tree) preSearch() {
    fmt.Println(tree.value)
    if tree.left != nil {
        tree.left.preSearch()
    }
    if tree.right != nil {
        tree.right.preSearch()
    }
}

func (tree Tree) centerSearch() {
    if tree.left != nil {
        tree.left.centerSearch()
    }
    fmt.Println(tree.value)
    if tree.right != nil {
        tree.right.centerSearch()
    }
}

func (tree Tree) afterSearch() {
    if tree.left != nil {
        tree.left.afterSearch()
    }
    if tree.right != nil {
        tree.right.afterSearch()
    }
    fmt.Println(tree.value)
}

func BFSSearch(node *Node) {
    queue := make([]*Node, 0) // slice 模拟先进先出队列
    queue = append(queue, node)
    result := make(map[int][]int, 0)
    height := 0
    for ;len(queue) > 0; {
        levelSize := len(queue)
        currentLevelItem := make([]int, 0)

        for i:=0; i<levelSize; i++{
            item := queue[0]
            queue = queue[1:]
            currentLevelItem = append(currentLevelItem, item.value)
            if item.leftChild != nil {
                queue = append(queue, item.leftChild)
            }
            if item.rightChild != nil {
                queue = append(queue, item.rightChild)
            }
        }
        result[height] = currentLevelItem
        height++
    }
    fmt.Println(result)
}

func DFSSearch(node *Node) (int) {
    max := 0
    return recursiveDFS(node, 0, max)
}


func recursiveDFS(node *Node, level int, max int) (int) {
    if node == nil {
        return 0
    }
    if node.leftChild == nil && node.rightChild == nil {
        if max < level {
            max = level
        }
        return max
    }
    return Max(recursiveDFS(node.leftChild, level + 1, max), recursiveDFS(node.rightChild, level + 1, max))
}

func main() {
    left := &Node{nil, nil, 9}
    right := &Node{&Node{&Node{nil, nil, 15}, nil, 15}, &Node{nil, nil, 7}, 20}
    root := &Node{left, right, 3}
    fmt.Println(DFSSearch(root))
}
