package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	element := l.Head
	for element != nil {
		element = l.Head.Next
	}

	return element.Data
}
