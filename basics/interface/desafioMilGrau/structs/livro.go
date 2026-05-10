package structs

type Livro struct {
	Nome string
	Preco float64
}

func (l *Livro) PrecoTotal() float64 {
	return l.Preco
}