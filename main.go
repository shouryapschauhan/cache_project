package main

import "fmt"

const SIZE = 5

type Node struct {
	value string
	left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.left = head

	return Queue{Head: head, Tail: tail}
}

type Hash map[string]*Node

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{value: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.value)
	left := n.left
	right := n.Right

	right.left = left
	left.Right = right

	c.Queue.Length -= 1
	delete(c.Hash, n.value)
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Add: %s", n.value)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n

	n.left = c.Queue.Head
	n.Right = tmp
	tmp.left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.value)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("Starting cache")

	cache := NewCache()

	for _, word := range []string{"parrot", "avocado", "dragon", "fruit", "tree", "potato", "tree"} {
		cache.Check(word)
		cache.Display()
	}
}
