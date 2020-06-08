package composite

import "fmt"

type Graphic interface {
	Move(x,y int)
	Draw()
}


type Dot struct{
	x, y int
}

func (d Dot) Move(x, y int) {
	d.x += x
	d.y += y
}

func (d Dot) Draw() {
	fmt.Printf("Draw Dot %d, %d\n", d.x , d.y)
}

type Circle struct {
	x, y int
	radius float32
}

func (c Circle) Move(x, y int) {
	c.x += x
	c.y += y
}

func (c Circle) Draw() {
	fmt.Printf("Draw Circle %d, %d, radius:%2f\n", c.x , c.y, c.radius)
}

type CompoundGraphic struct {
	childs []Graphic
}

func (c *CompoundGraphic) Add(g Graphic) {
	c.childs = append(c.childs, g)
}
func (c CompoundGraphic) Move(x, y int) {
	for _, g := range c.childs {
		g.Move(x, y)
	}
}

func (c CompoundGraphic) Draw() {
	for _, g := range c.childs {
		g.Draw()
	}
}


func PlayingAroundWithCompositePattern() {
	subBox := CompoundGraphic{}
	subBox.Add(Dot{1,2})
	subBox.Add(Circle{1,2, 23.0})

	allBox := CompoundGraphic{}
	allBox.Add(Dot{4,7})
	allBox.Add(subBox)
	allBox.Draw()
}