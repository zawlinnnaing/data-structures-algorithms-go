package channel

import (
	"bytes"
	"fmt"
)

var freeList = make(chan *bytes.Buffer, 100)
var serverChan = make(chan *bytes.Buffer)

var counter = 0

func load(bytes *bytes.Buffer) {
	bytes.Write([]byte{1, 2})
}

func client() {
	var b *bytes.Buffer
	select {
	case b = <-freeList:
	default:
		b = new(bytes.Buffer)
	}
	load(b)
	serverChan <- b
}

func process(receivedBuffer *bytes.Buffer) {
	counter += 1
}

func server() {
	for {
		b := <-serverChan
		process(b)
		select {
		case freeList <- b:
		default:
		}
	}
}

func RunLeakyBufferExample() {
	numOfClients := 10000
	go server()
	for i := 0; i < numOfClients; i++ {
		go client()
	}
	fmt.Printf("Actual number of clients %v. \n Number of requests server received %v.\n Number of packets dropped %v.\n", numOfClients, counter, numOfClients-counter)
}
