package toolset

import "testing"

func TestTools_RandomStringDifferentLength(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(15)
	if len(s) != 15 {
		t.Errorf("Expected string of length 15, got %d", len(s))
	}
}

func TestTools_RandomStringRandomness(t *testing.T) {
	var testTools Tools
	s1 := testTools.RandomString(10)
	s2 := testTools.RandomString(10)
	if s1 == s2 {
		t.Error("RandomString returned the same string twice")
	}
}
