package stringset

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := New("1")
	if s.Length() != 1 {
		t.Errorf("expected result to contain exactly one element, but got: %d", s.Length())
	}
	if !s.Contains("1") {
		t.Errorf("expected result to contain %#v, but was missing", "1")
	}

	s = New("1", "2", "3")
	if s.Length() != 3 {
		t.Errorf("expected result to contain exactly three elements, but got: %d", s.Length())
	}
	if !s.ContainsAll("1", "2", "3") {
		t.Error("expected element was missing")
	}

	s = New()
	if s.Length() != 0 {
		t.Fatalf("expected result to contain no elements, but got: %d", s.Length())
	}
}

func TestUnion(t *testing.T) {
	a := New("1", "2", "3", "4")
	b := New("5", "6", "7", "8")

	c := a.Union(b)

	if c.Length() != 8 {
		t.Fatalf("expected result to contain a total of 8 values after union operation, but got: %d", c.Length())
	}
	if !c.ContainsAll("1", "2", "3", "4", "6", "7", "8") {
		t.Error("expected element was missing")
	}
}

func TestIntersection(t *testing.T) {
	a := New("1", "2", "3", "4", "5", "6")
	b := New("2", "4", "6", "8")

	c := a.Intersection(b)

	if c.Length() != 3 {
		t.Fatalf("expected result to contain a total of 2 values after interesection operation, but got: %d", c.Length())
	}
	if !c.ContainsAll("2", "4", "6") {
		t.Error("expected element was missing")
	}
}

func TestDifference(t *testing.T) {
	a := New("1", "2", "3", "4", "5", "6", "7")
	b := New("6", "7")
	c := a.Difference(b)

	if c.Length() != 5 {
		t.Fatalf("expected result to contain a total of 3 values after difference operation, but got: %d", c.Length())
	}
	if !c.ContainsAll("1", "2", "3", "4", "5") {
		t.Error("expected element was missing")
	}
}
