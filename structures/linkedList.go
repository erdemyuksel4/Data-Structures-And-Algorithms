package structures

import "fmt"

type LinkedListNode struct {
	next  *LinkedListNode
	value int
}
type LinkedList struct {
	head *LinkedListNode
}

func (l *LinkedList) CreateList(node *LinkedListNode) {
	l.head = node
}
func CreateNode(data int, l *LinkedListNode) *LinkedListNode {
	return &LinkedListNode{
		value: data,
		next:  l,
	}
}
func (list *LinkedList) DeleteNode(l *LinkedListNode) {
	a := list.head
	for a != nil {
		if a.next.value == l.value {
			a.next = l.next
			break
		}
		a = a.next
	}
}
func (list *LinkedList) Next() {
	list.head = list.head.next
}
func (list *LinkedList) InsertNode(l *LinkedListNode, position int) {
	if position < 0 {
		fmt.Println("Invalid position")
		return
	}
	if position == 0 {
		l.next = list.head
		list.head = l
		return
	}
	a := list.head
	i := 1
	for a != nil {
		if i == position {
			l.next = a.next
			a.next = l
			return
		}
		a = a.next
		i++
	}
}
func PrintLinkedList(list *LinkedList) {
	l := list.head
	for l != nil {
		fmt.Printf("%d ", l.value)
		l = l.next
	}
}
