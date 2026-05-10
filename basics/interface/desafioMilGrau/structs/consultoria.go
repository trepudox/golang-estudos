package structs

type Consultoria struct {
	Nome      string
	Horas     int
	ValorHora float64
}

func (c *Consultoria) PrecoTotal() float64 {
	return c.ValorHora * float64(c.Horas)
}