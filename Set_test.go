package Set

import "testing"

func TestSetBasic(t *testing.T) {
	s := NewSet[string]()

	// Test Add
	s.Add("apple")
	s.Add("banana")

	if !s.Contains("apple") {
		t.Error("expected set to contain 'apple'")
	}

	if !s.Contains("banana") {
		t.Error("expected set to contain 'banana'")
	}

	if s.Contains("orange") {
		t.Error("did not expect 'orange' to be in the set")
	}

	// Test Size
	if s.Size() != 2 {
		t.Errorf("expected size 2, got %d", s.Size())
	}

	// Test Remove
	s.Remove("banana")
	if s.Contains("banana") {
		t.Error("expected 'banana' to be removed")
	}
	if s.Size() != 1 {
		t.Errorf("expected size 1 after removal, got %d", s.Size())
	}

	// Test Clear
	s.Clear()
	if s.Size() != 0 {
		t.Errorf("expected size 0 after Clear, got %d", s.Size())
	}
}
