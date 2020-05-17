package main

type Point struct {
	x, y int
}

func (p *Point) Set(x, y int) {
	p.x = x
	p.y = y
}

func AddItemUseIndex(list *[]*Point, index int) {
	(*list)[index] = &Point{}
}

func AddItemUseAppend(list *[]*Point) {
	*list = append(*list, &Point{})
}
