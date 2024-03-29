package toolset

import "testing"

func TestTools_RandomStringDifferentLength(t *testing.T) {
	var testTools Tools
	s := testTools.RandomPassword(15)
	if len(s) != 15 {
		t.Errorf("Expected string of length 15, got %d", len(s))
	}
}

func TestTools_RandomStringRandomness(t *testing.T) {
	var testTools Tools
	s1 := testTools.RandomPassword(10)
	s2 := testTools.RandomPassword(10)
	if s1 == s2 {
		t.Error("RandomPassword returned the same password twice")
	}
}
