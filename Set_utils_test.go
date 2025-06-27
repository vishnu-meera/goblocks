package Set

import "testing"

func TestUnionStatic(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)

	b := NewSet[int]()
	b.Add(2)
	b.Add(3)

	c := NewSet[int]()
	c.Add(4)

	result := Union(a, b, c)
	expected := []int{1, 2, 3, 4}
	for _, val := range expected {
		if !result.Contains(val) {
			t.Errorf("expected union to contain %d", val)
		}
	}
}

func TestIntersectionStatic(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewSet[int]()
	b.Add(2)
	b.Add(3)

	c := NewSet[int]()
	c.Add(3)
	c.Add(2)

	result := Intersection(a, b, c)
	if result.Size() != 2 || !result.Contains(2) || !result.Contains(3) {
		t.Errorf("expected intersection to contain 2 and 3 only")
	}
}
