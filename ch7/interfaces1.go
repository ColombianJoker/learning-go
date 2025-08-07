package main

import "fmt"

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

func main() {
	l := LinkedList{}
	l.Insert(int(0), int(1974))
	l.Insert(1, 2025)
	l.Insert(0, 1974)
	l.Insert(2, 2020)
	fmt.Println("l.0:", l.Value)
	fmt.Println("l.0.Next:", l.Next.Value)
	fmt.Println("l.0.Next.Next", l.Next.Next.Value)
}
