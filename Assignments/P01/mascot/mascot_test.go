package mascot_test

import (
	"testing"

	"github.com/SamOlatunde/4143-PLC/Assignments/P01/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
