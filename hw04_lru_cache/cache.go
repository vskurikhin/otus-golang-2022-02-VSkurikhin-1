package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache

	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := l.items[key]; ok {
		l.queue.MoveToFront(item)
		if i, ok := item.Value.(*cacheItem); ok {
			return i.value, true
		}
	}
	return nil, false
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := l.items[key]; ok {
		item.Value = &cacheItem{key: key, value: value}
		l.queue.MoveToFront(item)
		return true
	}
	item := l.queue.PushFront(&cacheItem{key: key, value: value})
	if l.queue.Len() > l.capacity {
		last := l.queue.Back()
		l.queue.Remove(last)
		if i, ok := last.Value.(cacheItem); ok {
			delete(l.items, i.key)
		}
	}
	l.items[key] = item
	return false
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}
