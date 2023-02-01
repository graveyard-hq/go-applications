package main

import (
   "fmt"
)

type Node struct {
    value int
    next  *Node
}

type Queue struct {
    head *Node
    tail *Node
}

func (q *Queue) Push(value int) {
    node := &Node{value: value}
    if q.tail == nil {
        q.head = node
        q.tail = node
        return
    }

    q.tail.next = node
    q.tail = node
}

func (q *Queue) Pop() int {
    if q.head == nil {
        return -1
    }

    value := q.head.value
    q.head = q.head.next
    if q.head == nil {
        q.tail = nil
    }

    return value
}

func main() {
    queue := &Queue{}

    queue.Push(1)
    queue.Push(2)
    queue.Push(3)

    fmt.Println(queue.Pop())
    fmt.Println(queue.Pop())
    fmt.Println(queue.Pop())
}
