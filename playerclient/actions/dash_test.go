package actions

import (
	"testing"
)

func TestDash(t *testing.T) {
	dashStr := Dash(2.743)
	if dashStr != "(dash 2.743)" {
		t.Fail()
	}
}
