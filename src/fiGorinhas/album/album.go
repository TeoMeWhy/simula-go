package album

import "math/rand"

type Album map[int]int

type Pacote []int

func NewPacote(n, N int) Pacote {
	pacote := Pacote{}
	for ; n > 0; n-- {
		pacote = append(pacote, rand.Intn(N)+1)
	}
	return pacote
}

func SimulaAlbum(n, N int) int {

	album := Album{}

	count := 0
	for len(album) < N {
		pacote := NewPacote(n, N)
		for _, v := range pacote {
			album[v]++
		}
		count++
	}

	return count
}

func SimulaAlbuns(n, N, ite int) int {

	soma := 0
	for i := ite; i > 0; i-- {
		count := SimulaAlbum(n, N)
		soma += count
	}

	return soma

}
