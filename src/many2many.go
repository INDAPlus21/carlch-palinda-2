package main

import (
  "fmt"
  "math/rand"
  "strconv"
  "sync"
  "time"
)

func main() {
  rand.Seed(time.Now().Unix())

  const strings = 32
  const producers = 4
  const consumers = 2

  before := time.Now()
  ch := make(chan string, producers)
  wgp := new(sync.WaitGroup)
  wgc := new(sync.WaitGroup)
  wgp.Add(producers)
  wgc.Add(consumers)
  for i := 0; i < producers; i++ {
    go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
  }
  for i := 0; i < consumers; i++ {
    go Consume("c"+strconv.Itoa(i), ch, wgc)
  }
  wgp.Wait()
  close(ch)
  wgc.Wait()
  fmt.Println("time:", time.Now().Sub(before))
}

func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
  for i := 0; i < n; i++ {
    RandomSleep(100)
    ch <- id + ":" + strconv.Itoa(i)
  }
  wg.Done()
}

func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
  for s := range ch {
    fmt.Println(id, "recieved", s)
    RandomSleep(100)
  }
  wg.Done()
}

func RandomSleep(n int) {
  time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
