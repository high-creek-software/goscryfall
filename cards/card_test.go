package cards

import "testing"

func TestBasicManaCost(t *testing.T) {
	card := Card{ManaCost: "{W}"}

	cost := card.ParseManaCost()
	expected := []string{"{W}"}

	if len(cost) != 1 {
		t.Errorf("expected a cost: got %d", len(cost))
	}

	set := cost[0]
	if len(set) != len(expected) {
		t.Errorf("returned length not the expected length: got %d", len(cost))
	}

	if set[0] != expected[0] {
		t.Errorf("expected %s got %s", expected[0], cost[0])
	}
}

func TestComplexManaCost(t *testing.T) {
	card := Card{ManaCost: "{1}{W}{U/P}"}
	cost := card.ParseManaCost()
	expected := []string{"{1}", "{W}", "{U/P}"}

	if len(cost) != 1 {
		t.Errorf("expected 1 cost set: got %d", len(cost))
	}

	set := cost[0]
	if len(set) != len(expected) {
		t.Errorf("cost not correct expected: %d got %d", len(expected), len(cost))
	}

	for idx := 0; idx < 3; idx++ {
		if set[idx] != expected[idx] {
			t.Errorf("cost item not correct expected: %s got %s", expected[idx], cost[idx])
		}
	}
}

func TestDualCosts(t *testing.T) {
	card := Card{ManaCost: "{1}{W}//{X}{W}{W}"}
	sets := card.ParseManaCost()

	if len(sets) != 2 {
		t.Errorf("expected 2 sets got: %d", len(sets))
	}

	firstExpected := []string{"{1}", "{W}"}
	secondExpected := []string{"{X}", "{W}", "{W}"}

	expected := [][]string{firstExpected, secondExpected}

	for set := 0; set < len(expected); set++ {
		for idx := 0; idx < len(expected[set]); idx++ {
			if sets[set][idx] != expected[set][idx] {
				t.Errorf("first set does not match got %s expected %s", sets[set][idx], expected[set][idx])
			}
		}
	}
}

func TestNoManaCost(t *testing.T) {
	card := Card{}
	cost := card.ParseManaCost()
	var expected []string

	if len(cost) != len(expected) {
		t.Errorf("cost doesn't meet expected")
	}
}
