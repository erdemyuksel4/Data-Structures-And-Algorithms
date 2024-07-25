package main

import (
	"fmt"
	"search"
	"sorting"
	"structures"
)

func main() {
	TestBinarySearchTree()
}

func TestBinarySearchTree() {
	var searchTree *structures.BinarySearchTree
	searchTree.NewBinarySearchTree(3)
	searchTree.Insert(5)
	fmt.Println(searchTree.Right().Value())
	searchTree.Insert(2)
	fmt.Println(searchTree.Left().Value())
	searchTree.Insert(6)
	fmt.Println(searchTree.Left().Value())
	searchTree.Insert(3)
	fmt.Println(searchTree.Right().Value())
	searchTree.Insert(7)
	fmt.Println(searchTree.Left().Value())
	searchTree.Insert(1)
	fmt.Println(searchTree.Right().Value())
	searchTree.Insert(15)
	fmt.Println(searchTree.Search(15))
}
func TestBinaryTree() {
	baseNode := structures.NewBinaryTreeNode(15)

	node2 := structures.NewBinaryTreeNode(22)
	node3 := structures.NewBinaryTreeNode(13)
	node4 := structures.NewBinaryTreeNode(14)
	node5 := structures.NewBinaryTreeNode(5)
	node6 := structures.NewBinaryTreeNode(62)
	node7 := structures.NewBinaryTreeNode(7)
	baseNode.Insert(node2)
	baseNode.Insert(node3)
	node2.Insert(node4)
	node2.Insert(node5)
	node3.Insert(node6)
	node3.Insert(node7)
	fmt.Println(baseNode.Left().Right().Value())

}
func TestQueue() {
	var myQueue structures.Queue

	myQueue.Enqueue('A')
	myQueue.Enqueue('B')
	myQueue.Enqueue('C')
	fmt.Println("Queue: ", myQueue.Queue)
	myQueue.Dequeue()
	fmt.Println("Dequeue: ", myQueue.Queue)

	fmt.Println("Peek: ", myQueue.Peek())

	fmt.Println("isEmpty: ", myQueue.IsEmpty())

	fmt.Println("Size: ", myQueue.Size())
}
func TestStack() {
	var myStack structures.Stack

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	fmt.Println("Stack: ", myStack.Stack)
	myStack.Pop()
	fmt.Println("Pop: ", myStack.Stack)

	fmt.Println("Peek: ", myStack.Peek())

	fmt.Println("isEmpty: ", myStack.IsEmpty())

	fmt.Println("Size: ", myStack.Size())
}
func TestLinkedList() {
	var list structures.LinkedList

	node4 := structures.CreateNode(4, nil)
	node3 := structures.CreateNode(3, node4)
	node2 := structures.CreateNode(2, node3)
	node1 := structures.CreateNode(1, node2)

	list.CreateList(node1)
	//list.DeleteNode(node3)

	node5 := structures.CreateNode(5, nil)

	list.InsertNode(node5, 5)
	structures.PrintLinkedList(&list)
}
func TestBinarySearch() {
	arr := search.BinarySearch([]int{2, 3, 7, 7, 11, 15, 25}, 11)
	fmt.Println(arr)
}
func TestBubbleSort() {
	arr := sorting.BubbleSort([]int{7, 12, 9, 11, 3})
	fmt.Println(arr)
}
func TestSelectionSort() {
	arr := []int{7, 12, 9, 11, 3}
	fmt.Println(arr)
	arr = sorting.SelectionSort(arr)
	fmt.Println(arr)
}
func TestQuickSort() {
	arr := []int{7, 3, 6, 2, 1, 1, 12, 9, 11, 3}
	fmt.Println(arr)
	arr = sorting.QuickSort(arr)
	fmt.Println(arr)
}
func TestCountingSort() {
	arr := []int{2, 3, 0, 2, 3, 2}
	fmt.Println(arr)
	arr = sorting.CountingSort(arr)
	fmt.Println(arr)

}
func TestRadixSort() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println(arr)
	arr = sorting.RadixSort(arr)
	fmt.Println(arr)
}
func TestInsertionSort() {
	arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println(arr)
	arr = sorting.InsertionSort(arr)
	fmt.Println(arr)
}
