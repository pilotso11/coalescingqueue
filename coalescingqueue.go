// MIT License
//
// Copyright (c) 2023 Seth Osher
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package coalescingqueue is a trivial generic implementation of a coalescing queue data structure in GO.
// A coalescing queue exhibits normal FIFO queue dynamics except that an insert of an item already in the queue
// does not increase the queue length.  A coalescing queue is a useful data structure in message distribution
// systems where there is the potential for multiple producers to produce duplicate results.  It is also useful
// for handling slow consumer problems where the data lends itself to coalescing.
package coalescingqueue

import (
	"sync"
)

// CoalescingQueue of T is a FIFO (first-in, first-out) queue of T
// that ensures duplicate items of T added to the queue are only added once.
// Duplicate items retain their order in the queue.
type CoalescingQueue[T comparable] struct {
	lock   sync.Mutex
	lookup map[T]bool
	queue  []T
}

func NewCoalescingQueue[T comparable]() *CoalescingQueue[T] {
	return &CoalescingQueue[T]{
		lookup: make(map[T]bool),
	}
}

// Push an item onto the queue.
// Returns true if the item was added and false if the item was already present.
func (q *CoalescingQueue[T]) Push(item T) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	_, ok := q.lookup[item]
	if ok {
		return false
	}
	q.queue = append(q.queue, item)
	q.lookup[item] = true
	return true
}

// Pop an item off the queue.
// Returns and item and true if the queue is not empty,
// otherwise and empty T and false
func (q *CoalescingQueue[T]) Pop() (item T, found bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) == 0 {
		found = false
		return
	}
	item = q.queue[0]
	q.queue = q.queue[1:]
	delete(q.lookup, item)
	found = true
	return
}

// Size returns the length of the queue
func (q *CoalescingQueue[T]) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.queue)
}

// Peek returns the head of the queue without removing it.
// If an item was found the item and true are returned.
// If the queue is empty and empty item and false are returned.
func (q *CoalescingQueue[T]) Peek() (item T, found bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) == 0 {
		found = false
		return
	}
	item = q.queue[0]
	found = true
	return
}
