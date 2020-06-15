package lexer

import (
	"reflect"
	"testing"
)

func TestSeeGlobal38(t *testing.T) {
	gPos, err := Look("(see_global 38 ((g r) 52.5 0) ((g l) -52.5 0) ((b) 7.26671 0.0139586 0.106379 -0.000523289) ((p \"ggb-lib-test\" 1) 4.40378 -5.68886 0.154905 -0.0812557 0 63))\x00")
	if err != nil {
		t.Fail()
	}
	if gPos == nil {
		t.FailNow()
	}

	expectedGPos := GlobalPositions{
		Time: 38,
		Objects: map[string][]string{
			"g r":                  {"52.5", "0"},
			"g l":                  {"-52.5", "0"},
			"b":                    {"7.26671", "0.0139586", "0.106379", "-0.000523289"},
			"p \"ggb-lib-test\" 1": {"4.40378", "-5.68886", "0.154905", "-0.0812557", "0", "63"},
		},
	}

	if !reflect.DeepEqual(expectedGPos, *gPos) {
		t.Fail()
	}
}

func TestOkLook100(t *testing.T) {
	gPos, err := Look("(ok look 100 ((g r) 52.5 0) ((g l) -52.5 0) ((b) 2.93611 0.0134586 0.940379 -0.00212289) ((p \"ggb-lib-test\" 1) -5.32903 -28.9635 -0.00217839 0.00301632 -173 -7))\x00")
	if err != nil {
		t.Fail()
	}
	if gPos == nil {
		t.FailNow()
	}

	expectedGPos := GlobalPositions{
		Time: 100,
		Objects: map[string][]string{
			"g r":                  {"52.5", "0"},
			"g l":                  {"-52.5", "0"},
			"b":                    {"2.93611", "0.0134586", "0.940379", "-0.00212289"},
			"p \"ggb-lib-test\" 1": {"-5.32903", "-28.9635", "-0.00217839", "0.00301632", "-173", "-7"},
		},
	}

	if !reflect.DeepEqual(expectedGPos, *gPos) {
		t.Fail()
	}
}

// (see_global 300 ((g r) 52.5 0) ((g l) -52.5 0) ((b) 0 0 0 0) ((p "ggb-lib-test" 1) 24.4006 -7.79101 0.0380215 0.172137 51 90))
// (ok look 300 ((g r) 52.5 0) ((g l) -52.5 0) ((b) 0 0 0 0) ((p "ggb-lib-test" 1) 24.4006 -7.79101 0.0380215 0.172137 51 90))
