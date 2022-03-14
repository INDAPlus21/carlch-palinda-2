package main

import (
  "fmt"

  "time"

  "sync"
)

/*
Code does not work as it finished before the print goroutine is finished which means not all elements are printed.
*/

var wg sync.WaitGroup   // Mkae WaitGroup
func main() {
  ch := make(chan int)
  wg.Add(1)             // Add one to the waitgroup
  go Print(ch)
  for i := 1; i <= 11; i++ {
    ch <- i
  }
  close(ch)
  wg.Wait()             // Wait for waitgroup to be done
}

func Print(ch <-chan int) {
  for n := range ch {
    time.Sleep(10 * time.Millisecond)
    fmt.Println(n)
  }
  defer wg.Done()       // finish waitgroup allowing program to exit
}
