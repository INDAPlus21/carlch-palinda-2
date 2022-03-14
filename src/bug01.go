package main

import (
	"fmt"
)

func main() {
	/*
	  The code is waiting for a channel but no goroutine has access to the channel which means the program is waiting for something which never will happen.

	  ch := make(chan string)
	  ch <- "Hello world!"
	  fmt.Println(<-ch)
	*/

	/*
	  This code resolves the deadlock as it contains a goroutine with access to the channel. The code is not waiting for something which will not appear no longer and therefore is not considered deadlocked by the compiler.
	*/
	ch := make(chan string)
	go func() {
		ch <- "Hello world!"
		close(ch)
	}()
	fmt.Println(<-ch)
}
