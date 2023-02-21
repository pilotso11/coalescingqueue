package main

import (
	"fmt"

	"github.com/pilotso11/coalescingqueue"
)

func main() {
	queue := coalescingqueue.NewCoalescingQueue[string]()

	queue.Push("item1")
	queue.Push("item2")
	queue.Push("item1") // does nothing
	queue.Push("item3")
	queue.Push("item3") // does nothing
	queue.Push("item2") // does nothing

	for queue.Size() > 0 {
		item, _ := queue.Pop()
		fmt.Println(item)
	}
}

// Expected Output:
// item1
// item2
// item3
