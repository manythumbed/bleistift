package bleistift

func project(a, b point, d float32) point {
	return a.plus(b.minus(a).scale(d))
}

type point struct {
	X, Y float32
}

func (p point) plus(p1 point) point {
	return point{p.X + p1.X, p.Y + p1.Y}
}

func (p point) minus(p1 point) point {
	return point{p.X - p1.X, p.Y - p1.Y}
}

func (p point) scale(s float32) point {
	return point{s * p.X, s * p.Y}
}

type construction struct {
	Points map[string]point
}

func (c *construction) define(name string, value point) {
	c.Points[name] = value
}

type instruction struct {
}

type renderer interface {
	render(i instruction)
}

func construct(c construction, instructions []instruction, r renderer) {

}
