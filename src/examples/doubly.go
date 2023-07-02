package examples

import (
	"fmt"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/linkedlists/doubly"
)

func RunDoublyLinkedListExample() {
	doubly := doubly.NewDoublyLinkedList[string]()

	doubly.Prepend("hello")
	doubly.Append("World")
	doubly.InsertAt("Zhiyu", 2)
	doubly.RemoveAt(2)
	fmt.Println("Result:", doubly)
	doubly.Append("Zhiyu")
	doubly.Remove("Zhiyu")
	fmt.Println("After removed", doubly)
}
