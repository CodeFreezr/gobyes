package structures

// snode is single linked list node
type snode struct {
	data interface{}
	next *snode
}

type Item struct {
	Value interface{}
	Next  *Item
	Prev  *Item
}

type LinkedList struct {
	head *Item
	tail *Item
}

func (l *LinkedList) Prepend(item *Item) *Item {
	head := l.head
	l.head = item
	item.Next = head

	if head != nil {
		head.Prev = item
	}

	if l.tail == nil {
		l.tail = item
	}

	return item
}

func (l *LinkedList) Append(item *Item) *Item {
	tail := l.tail
	l.tail = item

	if tail != nil {
		tail.Next = item
	}

	if l.head == nil {
		l.head = item
	}

	return item
}

func (l *LinkedList) Shift() *Item {
	head := l.head

	if head != nil {
		l.head = head.Next
	}

	if l.head == nil {
		l.tail = nil
	}

	return head
}

func (l *LinkedList) Head() *Item {
	return l.head
}

func (l *LinkedList) Insert(before *Item, item *Item) *Item {
	prev := before.Prev
	before.Prev = item
	if prev != nil {
		prev.Next = item
	}

	return item
}

func (l *LinkedList) Remove(item *Item) {
	if item == nil {
		return
	}

	if item == l.head && item == l.tail {
		l.head = nil
		l.tail = nil
	} else if item == l.head {
		l.head = l.head.Next
		l.head.Prev = nil
	} else if item == l.tail {
		l.tail = l.tail.Prev
		l.tail.Next = nil
	} else {
		item.Prev.Next, item.Next.Prev = item.Next, item.Prev
	}
}

func NewLinkedList(values ...interface{}) *LinkedList {
	list := &LinkedList{}

	if len(values) == 0 {
		return list
	}
	prev := &Item{values[0], nil, nil}
	list.head = prev
	list.tail = prev

	for i := 1; i < len(values); i++ {
		item := &Item{values[i], nil, nil}
		prev.Next = item
		item.Prev = prev
		prev = item
		list.tail = item
	}
	return list
}
