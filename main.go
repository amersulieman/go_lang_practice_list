package main

import "fmt"

func main() {
	mylist := LinkedList{}
	mylist.delete(5)
	mylist.append("Amer")
	mylist.append("5")
	mylist.append(10)
	mylist.append("sulieman")
	fmt.Println(mylist)
	mylist.delete(10)
	fmt.Println(mylist)
}
