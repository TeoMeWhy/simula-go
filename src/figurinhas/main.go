package main

import (
	"figurinhas/album"
	"log"
	"sync"
	"time"
)

func worker(n, N, ite int, ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	ch <- album.SimulaAlbuns(n, N, ite)
	log.Printf("\nEncerrando a thread %d...", i)
}

func listener(ch chan int, ite int) {
	soma := 0
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				log.Println(soma / ite)
				return
			}
			soma += v
		}
	}
}

func main() {

	n := 5
	N := 550
	ite := 1000000
	ch := make(chan int, ite)
	threads := 5000
	partition := ite / threads

	wg := sync.WaitGroup{}

	go listener(ch, ite)

	for ; threads > 0; threads-- {
		wg.Add(1)
		go worker(n, N, partition, ch, &wg, threads)
	}

	wg.Wait()
	close(ch)
	time.Sleep(time.Millisecond * 10)
}
