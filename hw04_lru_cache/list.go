package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List
	size  int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(value interface{}) *ListItem {
	return l.pushFront(&ListItem{
		Value: value,
		Next:  l.front,
	})
}

func (l *list) pushFront(item *ListItem) *ListItem {
	if l.front != nil {
		l.front.Prev = item
		l.front = item
	} else {
		l.front = item
		l.back = item
	}
	l.size++
	return item
}

func (l *list) PushBack(value interface{}) *ListItem {
	return l.pushBack(&ListItem{
		Value: value,
		Prev:  l.back,
	})
}

func (l *list) pushBack(item *ListItem) *ListItem {
	if l.back != nil {
		l.back.Next = item
		l.back = item
	} else {
		l.front = item
		l.back = item
	}
	l.size++
	return item
}

func (l *list) Remove(item *ListItem) {
	if l.size > 0 {
		done := false
		prev := item.Prev
		next := item.Next
		if l.front == item {
			l.front = item.Next
			done = true // обработали случай когда удалили первый в списке
		}
		if l.back == item {
			l.back = item.Prev
			done = true // обработали случай когда удалили последний в списке
		}
		if prev != nil {
			prev.Next = next
			done = true // обработали случай для предыдущего, когда где-то между первым и вторым
		}
		if next != nil {
			next.Prev = prev
			done = true // обработали случай для следующего, когда где-то между первым и вторым
		}
		if done {
			item.Next = nil // почистили ссылки в
			item.Prev = nil // удаляемом элементе
			l.size--
		}
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	i.Next = l.front
	l.pushFront(i)
}
