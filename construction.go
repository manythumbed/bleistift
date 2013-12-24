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
	points map[string]point
}

func (c *construction) define(name string, value point) {
	c.points[name] = value
}

func (c *construction) point(name string) (point, bool) {
	p, ok := c.points[name]

	return p, ok
}

type renderer interface {
	line(p1, p2 point)
	curve(p1, p2, p3 point)
}

type curve struct {
	p1, p2, p3 string
}

type line struct {
	p1, p2 string
}

func construct(c construction, instructions []interface{}, r renderer) {
	for _, v := range instructions {
		switch t := v.(type) {
		default:
		case curve:
			p1, ok := c.point(t.p1)
			if !ok {
			}
			p2, ok := c.point(t.p2)
			if !ok {
			}
			p3, ok := c.point(t.p3)
			if !ok {
			}
			r.curve(p1, p2, p3)
		}

	}
}
