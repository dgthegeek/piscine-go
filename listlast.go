package piscine

func ListLast(l *List) interface{} {
	element := l.Head
	for element != nil {
		element = element.Next
	}

	if l.Head == nil {
		return nil
	}

	return element.Data
}
