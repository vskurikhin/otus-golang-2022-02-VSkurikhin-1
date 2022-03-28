package hw04lrucache

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestListFront(t *testing.T) {
	l := NewList()
	require.Nil(t, l.Front())
}

func TestListBack(t *testing.T) {
	l := NewList()
	require.Nil(t, l.Back())
}

func TestListLenEmptyList(t *testing.T) {
	l := NewList()
	require.Equal(t, 0, l.Len())
}

func TestListPushBack(t *testing.T) {
	l := NewList()
	l.PushBack("test1")
	require.Equal(t, 1, l.Len())
	require.Equal(t, "test1", l.Front().Value)
	require.Equal(t, "test1", l.Back().Value)
	require.Nil(t, l.Front().Prev)
	require.Nil(t, l.Front().Next)
	require.Nil(t, l.Back().Next)
	require.Nil(t, l.Back().Prev)
	l.PushBack("test2")
	require.Equal(t, 2, l.Len())
	require.Equal(t, "test1", l.Front().Value)
	require.Equal(t, "test2", l.Back().Value)
	require.Equal(t, l.Front().Next, l.Back())
	require.Equal(t, l.Front(), l.Back().Prev)
	require.Nil(t, l.Front().Prev)
	require.Nil(t, l.Back().Next)
}

func TestListRemoveOne(t *testing.T) {
	l := NewList()
	require.Equal(t, 0, l.Len())
	l.Remove(nil)
	require.Equal(t, 0, l.Len())
	test1 := l.PushBack("test1")
	l.Remove(test1)
	require.Nil(t, test1.Next)
	require.Nil(t, test1.Prev)
	require.Equal(t, 0, l.Len())
	require.Nil(t, l.Front())
	require.Nil(t, l.Back())
}

func TestListRemoveTwo(t *testing.T) {
	l := NewList()
	test1 := l.PushBack("test1")
	test2 := l.PushBack("test2")
	require.Equal(t, 2, l.Len())
	l.Remove(test2)
	require.Nil(t, test2.Next)
	require.Nil(t, test2.Prev)
	require.Equal(t, 1, l.Len())
	require.Equal(t, l.Front(), l.Back())
	require.Nil(t, l.Front().Prev)
	require.Nil(t, l.Back().Next)
	l.Remove(test1)
	require.Equal(t, 0, l.Len())
	require.Nil(t, l.Front())
	require.Nil(t, l.Back())
	l.Remove(test2)
	l.Remove(test1)
	require.Equal(t, 0, l.Len())
	require.Nil(t, l.Front())
	require.Nil(t, l.Back())
}

func TestListMoveToFront(t *testing.T) {
	l := NewList()
	l.PushBack("test1")
	test2 := l.PushBack("test2")
	l.MoveToFront(test2)
	require.Equal(t, 2, l.Len())
	require.Equal(t, "test2", l.Front().Value)
	require.Equal(t, "test1", l.Back().Value)
	require.Equal(t, l.Front().Next, l.Back())
	require.Equal(t, l.Front(), l.Back().Prev)
	require.Nil(t, l.Front().Prev)
	require.Nil(t, l.Back().Next)
}
