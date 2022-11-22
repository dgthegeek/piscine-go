package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	element := l.Head
	for element.Next != nil {
		element = element.Next
	}

	return element.Data
}
