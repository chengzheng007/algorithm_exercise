package main

// reverse a linked list from head to tail
func reverseList(head *Node) {
	// no element or has one element, no need to process
	if head == nil || head.Next == nil {
		return
	}
	p := head.Next
	q := p.Next
	for q != nil {
		r := q.Next
		q.Next = p
		p = q
		q = r
	}
	head.Next.Next = nil
	head.Next = p
}
