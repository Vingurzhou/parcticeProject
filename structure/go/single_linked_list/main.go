package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func (this *ListNode) Insert(ln *ListNode) {
	for this.next != nil {
		this = this.next
	}
	this.next = ln
}
func (this *ListNode) Show() {
	for this.next != nil {
		fmt.Print(this.val, "\t")
		this = this.next
	}
	if this.next == nil {
		fmt.Println(this)
	}

}
func (this *ListNode) Delete(val int) {
	for this.next != nil {
		if this.next.val == val {
			this.next = this.next.next
		}
		this = this.next
	}
	if this.next == nil {
	}
}
func main() {
	head := &ListNode{}
	head.Show()
	head.Insert(&ListNode{val: 1})
	head.Insert(&ListNode{val: 2})
	head.Insert(&ListNode{val: 3})
	head.Show()
	head.Delete(2)
	head.Show()
}
