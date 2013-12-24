package bleistift

import (
	"fmt"
)

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

func new() construction {
	return construction{map[string]point{}}
}
func (c *construction) define(name string, value point) {
	c.points[name] = value
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

type errors struct {
	list []string
}

func (es errors) Error() string {
	return fmt.Sprintf("There are %d errors", len(es.list))
}

func (es errors) add(e error) {
	if e != nil {
		es.list = append(es.list, e.Error())
	}
}

func construct(c construction, instructions []interface{}, r renderer) (bool, error) {
	e := errors{}
	for _, v := range instructions {
		switch t := v.(type) {
		default:
		case curve:
			p1, err := c.point(t.p1)
			e.add(err)
			p2, err := c.point(t.p2)
			e.add(err)
			p3, err := c.point(t.p3)
			e.add(err)
			r.curve(p1, p2, p3)
		}
	}

	return true, e
}

func (c *construction) point(name string) (point, error) {
	p, ok := c.points[name]
	if ok {
		return p, nil
	}

	return point{}, fmt.Errorf("Unable to find point (%s) in construction", name)
}
