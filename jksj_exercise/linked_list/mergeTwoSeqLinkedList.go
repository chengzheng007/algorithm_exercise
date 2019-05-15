package main

// merge two sequential linked list
func mergeTwoSeqLinkedList(h1 *Node, h2 *Node) *Node {
	if h1 == nil || h1.Next == nil {
		return h2
	} else if h2 == nil || h2.Next == nil {
		return h1
	}
	p := h1.Next
	q := h2.Next
	newh := new(Node)
	r := newh
	for p != nil && q != nil {
		if p.Data < q.Data {
			r.Next = p
			p = p.Next
		} else {
			r.Next = q
			q = q.Next
		}
		r = r.Next
	}
	// only one linked list has reaining element
	if p != nil {
		r.Next = p
	}
	if q != nil {
		r.Next = q
	}
	return newh
}
