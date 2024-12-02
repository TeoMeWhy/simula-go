package main

import (
	"fiGorinhas/album"
	"flag"
	"log"
	"sync"
	"time"
)

func worker(n, N, ite int, ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	ch <- album.SimulaAlbuns(n, N, ite)
	log.Printf("Encerrando a thread %d...", i)
}

func listener(ch chan int, ite int) {
	soma := 0
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				log.Println("\n\nMédia para completar album:", soma/ite)
				return
			}
			soma += v
		}
	}
}

func main() {

	n := flag.Int("n", 5, "Tamanho de cada pacote de figurinhas, i.e. quantas figurinahs em cada pacote.")
	N := flag.Int("N", 550, "Tamanho do álbum, i.e. quantas figurinhas são necessárias para completar o álbum.")
	ite := flag.Int("ite", 1000, "Quantidade de albuns gerados na simulação.")
	threads := flag.Int("threads", 1, "Quantidade de processos na simulação.")
	flag.Parse()

	ch := make(chan int, *ite)
	partition := *ite / *threads

	wg := sync.WaitGroup{}

	go listener(ch, *ite)

	for ; *threads > 0; *threads-- {
		wg.Add(1)
		go worker(*n, *N, partition, ch, &wg, *threads)
	}

	wg.Wait()
	close(ch)
	time.Sleep(time.Millisecond * 10)
}
