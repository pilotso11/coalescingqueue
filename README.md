# A trivial generic coalescing queue data strucgture for go/golang

A coalescing queue exhibits normal queue semantics except that duplicate additions
of items already in the queue does not extend the queue length.

A useful data structure for messaging applications where there are either multiple producers
that may produce duplicates or slow consumers.

## Installation

`go get github.com/pilotso11/coalescingqueue`

## Example 
```go
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
```
On GO Playground https://go.dev/play/p/SJHxfD8hLnq
