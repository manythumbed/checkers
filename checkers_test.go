package checkers

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type S struct {

}

var _ = Suite(&S{})

func (s *S) TestToleranceEquality(c *C) {
	c.Check(1.0, EqualsWithTolerance, 1.25, 0.5)
	c.Check(1.0, Not(EqualsWithTolerance), 1.25, 0.05)
}

func (s *S) TestBounds(c *C) {
	c.Check(1.0, Between, 0.0, 1.5)
	c.Check(1.0, Not(Between), 2.0, 2.5)
}

type x struct {
	Val string
}

type y struct {
	Val int
}

func (s *S) TestContainsType(c *C) {
	a := []int{2, 3, 4}
	c.Check(containsType(a, a[0]), Equals, true)
	c.Check(containsType(a, a), Equals, false)
	c.Check(containsType(a, "x"), Equals, false)

	b := []x{x{"1"}, x{"2"}}
	c.Check(containsType(b, b[0]), Equals, true)
	c.Check(containsType(b, y{0}), Equals, false)
}

func (s *S) TestContains(c *C) {
	a := []int{2, 3, 4}
	c.Check(a, Contains, a[0])
	c.Check(a, Contains, a[1])
	c.Check(a, Contains, a[2])
	c.Check(a, Contains, 2)
	c.Check(a, Contains, 3)
	c.Check(a, Contains, 4)
	c.Check(a, Not(Contains), 5)
	c.Check(a, Not(Contains), a)
	c.Check(a, Not(Contains), "x")

	b := []x{x{"1"}, x{"2"}}
	c.Check(b, Contains, b[0])
	c.Check(b, Contains, b[1])
	c.Check(b, Contains, x{"1"})
	c.Check(b, Contains, x{"2"})
	c.Check(b, Not(Contains), x{"3"})
	c.Check(b, Not(Contains), y{0})
}
