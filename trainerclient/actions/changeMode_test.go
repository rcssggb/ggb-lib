package actions

import (
	"testing"

	"github.com/rcssggb/ggb-lib/rcsscommon"
)

func TestChangeModeBeforeKickOff(t *testing.T) {
	cmStr := ChangeMode(rcsscommon.PlayModeBeforeKickOff)
	if cmStr != "(change_mode before_kick_off)" {
		t.Fail()
	}
}

// TODO: add more unit tests for the ChangeMode functions (not really a priority but would be nice)
