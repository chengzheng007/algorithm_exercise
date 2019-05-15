package main

// find the mid node of a linked list
func middleOfLinkedList(head *Node) *Node {
	singleP := head
	doubleP := singleP
	if singleP == nil {
		return singleP
	}
	for doubleP != nil && doubleP.Next != nil {
		singleP = singleP.Next
		doubleP = doubleP.Next.Next
	}
	return singleP
}
