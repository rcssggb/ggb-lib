package parser

import (
	"reflect"
	"testing"

	"github.com/rcssggb/ggb-lib/trainerclient/lexer"
)

func TestSeeGlobal38(t *testing.T) {
	errCh := make(chan string, 64)

	gpSymbols := lexer.GlobalPositions{
		Time: 38,
		Objects: map[string][]string{
			"g r":                  {"52.5", "0"},
			"g l":                  {"-52.5", "0"},
			"b":                    {"7.26671", "0.0139586", "0.106379", "-0.000523289"},
			"p \"ggb-lib-test\" 1": {"4.40378", "-5.68886", "0.154905", "-0.0812557", "0", "63"},
		},
	}

	gPos := Look(gpSymbols, errCh)
	close(errCh)

	if len(errCh) != 0 {
		t.Fail()
	}
	if gPos == nil {
		t.FailNow()
	}

	expectedGPos := GlobalPositions{
		Time: 38,
		Ball: AbsPosition{
			X:      7.26671,
			Y:      0.0139586,
			DeltaX: 0.106379,
			DeltaY: -0.000523289,
		},
		GoalLeft: AbsPosition{
			X: -52.5,
			Y: 0,
		},
		GoalRight: AbsPosition{
			X: 52.5,
			Y: 0,
		},
		Teams: map[string]Team{
			"ggb-lib-test": {
				1: AbsPosition{
					X:         4.40378,
					Y:         -5.68886,
					DeltaX:    0.154905,
					DeltaY:    -0.0812557,
					BodyAngle: 0,
					NeckAngle: 63,
				},
			},
		},
	}

	if !reflect.DeepEqual(expectedGPos, *gPos) {
		t.Fail()
	}
}

func TestOkLook100(t *testing.T) {
	errCh := make(chan string, 64)

	gpSymbols := lexer.GlobalPositions{
		Time: 100,
		Objects: map[string][]string{
			"g r":                  {"52.5", "0"},
			"g l":                  {"-52.5", "0"},
			"b":                    {"2.93611", "0.0134586", "0.940379", "-0.00212289"},
			"p \"ggb-lib-test\" 1": {"-5.32903", "-28.9635", "-0.00217839", "0.00301632", "-173", "-7"},
		},
	}

	gPos := Look(gpSymbols, errCh)
	close(errCh)

	if len(errCh) != 0 {
		t.Fail()
	}
	if gPos == nil {
		t.FailNow()
	}

	expectedGPos := GlobalPositions{
		Time: 100,
		Ball: AbsPosition{
			X:      2.93611,
			Y:      0.0134586,
			DeltaX: 0.940379,
			DeltaY: -0.00212289,
		},
		GoalLeft: AbsPosition{
			X: -52.5,
			Y: 0,
		},
		GoalRight: AbsPosition{
			X: 52.5,
			Y: 0,
		},
		Teams: map[string]Team{
			"ggb-lib-test": {
				1: AbsPosition{
					X:         -5.32903,
					Y:         -28.9635,
					DeltaX:    -0.00217839,
					DeltaY:    0.00301632,
					BodyAngle: -173,
					NeckAngle: -7,
				},
			},
		},
	}

	if !reflect.DeepEqual(expectedGPos, *gPos) {
		t.Fail()
	}
}

// (see_global 300 ((g r) 52.5 0) ((g l) -52.5 0) ((b) 0 0 0 0) ((p "ggb-lib-test" 1) 24.4006 -7.79101 0.0380215 0.172137 51 90))
// (ok look 300 ((g r) 52.5 0) ((g l) -52.5 0) ((b) 0 0 0 0) ((p "ggb-lib-test" 1) 24.4006 -7.79101 0.0380215 0.172137 51 90))
