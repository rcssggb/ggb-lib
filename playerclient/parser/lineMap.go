package parser

import "github.com/rcssggb/ggb-lib/rcsscommon"

var lineMap = map[string]rcsscommon.LineID{
	"l t": rcsscommon.LineTop,
	"l r": rcsscommon.LineRight,
	"l b": rcsscommon.LineBottom,
	"l l": rcsscommon.LineLeft,
}
