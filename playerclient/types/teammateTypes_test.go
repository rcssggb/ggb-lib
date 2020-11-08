package types

import (
	"reflect"
	"testing"
)

func TestDefaultTeammateTypes(t *testing.T) {
	tmt := DefaultTeammateTypes()

	tmtExpected := TeammateTypes{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0}

	if !reflect.DeepEqual(tmt, tmtExpected) {
		t.Fail()
	}
}
