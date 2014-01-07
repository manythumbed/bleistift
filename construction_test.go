package bleistift

import (
	"fmt"
	"testing"
)

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

	log := &logRenderer{}
	err := construct(c, []interface{}{line{"a", "a"}, curve{"a", "a", "a"}}, log)
	if err != nil {
		t.Errorf("Should have constructed without errors. %v", err)
	}

	if len(log.instructions) != 2 {
		t.Errorf("Expected 2 instructions, received %d", len(log.instructions))
	}
}

func TestConstructionWithErrors(t *testing.T) {
	c := new()
	c.define("a", point{1, 1})

	log := &logRenderer{}
	err := construct(c, []interface{}{line{"b", "b"}, curve{"b", "b", "b"}}, log)
	if err == nil {
		t.Errorf("Should have returned errors")
	}
	if err.Error() != "There are 5 errors" {
		t.Errorf("%v", err.Error())
	}
	if len(log.instructions) != 2 {
		t.Errorf("Expected 2 instructions, received %d", len(log.instructions))
	}
}

func TestHollowRectangle(t *testing.T) {
	construction := new()
	construction.define("a", point{1, 1})
	construction.define("b", point{4, 1})
	construction.define("c", point{4, 4})
	construction.define("d", point{1, 4})

	log := &logRenderer{}
	lineAB := line{"a", "b"}
	lineBC := line{"b", "c"}
	lineCD := line{"c", "d"}
	lineDA := line{"d", "a"}

	err := construct(construction, []interface{}{lineAB, lineBC, lineCD, lineDA}, log)
	if err != nil {
		t.Errorf("Should have rendered without errors")
	}

	if len(log.instructions) != 4 {
		t.Errorf("Should have rendered four instructions")
	}
}

type logRenderer struct {
	instructions []string
}

func (l *logRenderer) line(p1, p2 point) {
	l.instructions = append(l.instructions, fmt.Sprintf("L [%v %v]", p1, p2))
}

func (l *logRenderer) curve(p1, p2, p3 point) {
	l.instructions = append(l.instructions, fmt.Sprintf("C [%v %v %v]", p1, p2, p3))
}

func (l *logRenderer) move(p point) {
	l.instructions = append(l.instructions, fmt.Sprintf("M [%v]", p))
}
