package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Data     interface{}
	NextNode *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (ll *LinkedList) append(data interface{}) {

	newNode := Node{Data: data}
	if ll.Head == nil {
		ll.Head = &newNode
		ll.Tail = ll.Head
	} else {
		ll.Tail.NextNode = &newNode
		ll.Tail = &newNode
	}
	ll.Size++
}

func (ll *LinkedList) delete(data interface{}) {
	currentNode := ll.Head
	if ll.Size == 0 {
		println("Empty List")
		return
	}
	if reflect.DeepEqual(ll.Head.Data, data) {
		ll.Head = ll.Head.NextNode
		ll.Size--
		return
	}
	for currentNode.NextNode != nil {
		if reflect.DeepEqual(currentNode.NextNode.Data, data) {
			currentNode.NextNode = currentNode.NextNode.NextNode
			ll.Size--
			return
		}
		currentNode = currentNode.NextNode
	}
	println("Value not in list")
}

func (ll LinkedList) String() string {

	if ll.Size == 0 {
		return "Empty List"
	} else if ll.Size == 1 {
		return fmt.Sprint(ll.Head.Data)
	}
	currentNode := ll.Head
	list := fmt.Sprint(currentNode.Data)
	currentNode = currentNode.NextNode
	for currentNode != nil {
		list += fmt.Sprint(" --> ", currentNode.Data)
		currentNode = currentNode.NextNode
	}
	return list
}
