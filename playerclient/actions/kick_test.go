package actions

import (
	"testing"
)

func TestKick(t *testing.T) {
	kickStr := Kick(46, 0)
	if kickStr != "(kick 46.000 0.000)" {
		t.Fail()
	}
}
