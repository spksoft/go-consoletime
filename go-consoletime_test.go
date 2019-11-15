package gconsoletime

import "testing"

func TestMain(t *testing.T) {
	hasLogOne := hasLog("one")
	if hasLogOne == true {
		t.Errorf("It should have no log name one")
	}
	Log("one")
	hasLogOne = hasLog("one")
	if hasLogOne == false {
		t.Errorf("It should have log name one")
	}
	Log("one")
	hasLogOne = hasLog("one")
	if hasLogOne == true {
		t.Errorf("It should have no log name one")
	}
}
