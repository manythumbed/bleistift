package bleistift

import "testing"

func TestProjection(t *testing.T) {
	a := point{1.0, 1.0}
	b := point{2.0, 2.0}

	c := project(a, b, 1.0)
	pointsEqual(t, b, c)

	d := project(a, b, 0.0)
	pointsEqual(t, a, d)

	e := project(a, b, 0.5)
	pointsEqual(t, point{1.5, 1.5}, e)

	f := project(a, b, 1.5)
	pointsEqual(t, point{2.5, 2.5}, f)

	g := project(a, b, -1.0)
	pointsEqual(t, point{0, 0}, g)
}

func pointsEqual(t *testing.T, a, b point) {
	if a != b {
		t.Errorf("%v should equal %v)", a, b)
	}
}

func TestConstruction(t *testing.T) {
	c := new()
	c.define("a", point{1, 1})

	err := construct(c, []interface{}{line{"a", "a"}, curve{"a", "a", "a"}}, logRenderer{})
	if err != nil {
		t.Errorf("Should have constructed without errors. %v", err)
	}
}

func TestConstructionWithErrors(t *testing.T) {
	c := new()
	c.define("a", point{1, 1})

	err := construct(c, []interface{}{line{"b", "b"}, curve{"b", "b", "b"}}, logRenderer{})
	if err == nil {
		t.Errorf("Should have returned errors")
	}
	if err.Error() != "There are 3 errors" {
		t.Errorf("%v", err.Error())
	}
}

type logRenderer struct {
}

func (l logRenderer) line(p1, p2 point) {
}

func (l logRenderer) curve(p1, p2, p3 point) {
}
