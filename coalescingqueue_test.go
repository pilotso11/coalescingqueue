package coalescingqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoalescingQueue_Peek(t *testing.T) {
	q := NewCoalescingQueue[string]()
	assert.True(t, q.Push("new"))
	assert.True(t, q.Push("new2"))
	assert.True(t, q.Push("new3"))
	i, ok := q.Peek()
	assert.True(t, ok)
	assert.Equal(t, "new", i)
	i, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, "new", i)
	_, _ = q.Pop()
	i, ok = q.Peek()
	assert.True(t, ok)
	assert.Equal(t, "new2", i)
}

func TestCoalescingQueue_Pop(t *testing.T) {
	q := NewCoalescingQueue[string]()
	assert.True(t, q.Push("new"))
	assert.True(t, q.Push("new2"))
	assert.True(t, q.Push("new3"))
	i, ok := q.Pop()
	assert.True(t, ok)
	assert.Equal(t, "new", i)
	i, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, "new2", i)
	i, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, "new3", i)
	i, ok = q.Pop()
	assert.False(t, ok)
	assert.Equal(t, "", i)
	assert.True(t, q.Push("new"))
	i, ok = q.Pop()
	assert.True(t, ok)
	assert.Equal(t, "new", i)
}

func TestCoalescingQueue_Push(t *testing.T) {
	q := NewCoalescingQueue[string]()
	assert.True(t, q.Push("new"))
	assert.True(t, q.Push("new2"))
	assert.False(t, q.Push("new"))
	assert.True(t, q.Push("new3"))
	assert.False(t, q.Push("new2"))
	assert.False(t, q.Push("new2"))
	assert.True(t, q.Push(""))
	assert.False(t, q.Push(""))
}

func TestCoalescingQueue_Size(t *testing.T) {
	q := NewCoalescingQueue[string]()
	assert.True(t, q.Push("new"))
	assert.True(t, q.Push("new2"))
	assert.True(t, q.Push("new3"))
	assert.Equal(t, 3, q.Size())
}
