package main

import (
    . "launchpad.net/gocheck"
    "testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}
var _ = Suite(&S{})

func (s *S) TestProperDivisors(c *C) {
	expected := []int{1, 2, 4, 7, 14}
	c.Assert(properDivisors(28), Equals, expected)
}

func (s *S) TestSumReduce(c *C) {
	divisors := []int{1, 2, 4, 7, 14}
	c.Assert(sumReduce(divisors), Equals, 28)
}

func (s *S) TestIsAbundant(c *C) {
	c.Assert(isAbundant(28), Equals, false)
	c.Assert(isAbundant(12), Equals, true)
	c.Assert(isAbundant(11), Equals, false)
}
