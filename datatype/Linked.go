package datatype

func FlipLinked(linked *Node) *Node {

	p := linked.Next
	var q *Node
	linked.Next = nil
	for p != nil{
		pr := p.Next
		p.Next = q
		q = p
		p = pr
	}

	linked.Next = q
	return linked
}
