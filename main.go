package main

import (
	"fmt"

	"github.com/zawlinnnaing/data-structures-algorithms-go/src/linkedlists/singly"
)

func main() {
	fmt.Println("Running simple insert example:")
	singly.SimpleInsertExample()
	fmt.Println("Running InsertHeadExample:")
	singly.InsertHeadExample()
	fmt.Println("Running Insert After Example:")
	singly.InsertAfterExample()
	fmt.Println("Running Insert End example:")
	singly.InsertEndExample()
}