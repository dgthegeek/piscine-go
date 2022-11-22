package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}

	var precedent *NodeL
	var next *NodeL
	temp := l.Head
	current := l.Head

	for current != nil {
		next = current.Next
		current.Next = precedent
		precedent = current
		current = next
	}
	l.Head = precedent
	l.Tail = temp
	l.Tail.Next = nil
}
