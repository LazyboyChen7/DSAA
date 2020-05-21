package manachar

import (
	"testing"
)

func TestZeroManachar(t *testing.T) {
	s := ""
	re,l := Manachar(s)
	if re != "" || l != 0 {
		t.Error("Zero Test Error")
	}
}

func TestNormalSingleManachar(t *testing.T) {
	s := "abbahopxpo"
	re,l := Manachar(s)
	if re != "opxpo" || l != 5 {
		t.Error("Normal Test Error")
	}
}

func TestNormalDoubleManachar(t *testing.T) {
	s := "abbahopxxpoaavraats"
	re,l := Manachar(s)
	if re != "opxxpo" || l != 6 {
		t.Error("Normal Test Error")
	}
}