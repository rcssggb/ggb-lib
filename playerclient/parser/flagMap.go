package parser

import "github.com/rcssggb/ggb-lib/rcsscommon/flags"

var flagMap = map[string]flags.FlagID{
	"f c":       flags.CenterFlag,
	"f c t":     flags.CenterTopFlag,
	"f c b":     flags.CenterBotFlag,
	"g l":       flags.LeftGoalFlag,
	"f g l t":   flags.LeftGoalTopFlag,
	"f g l b":   flags.LeftGoalBotFlag,
	"g r":       flags.RightGoalFlag,
	"f g r t":   flags.RightGoalTopFlag,
	"f g r b":   flags.RightGoalBotFlag,
	"f p l c":   flags.LeftPenaltyCenterFlag,
	"f p l t":   flags.LeftPenaltyTopFlag,
	"f p l b":   flags.LeftPenaltyBotFlag,
	"f p r c":   flags.RightPenaltyCenterFlag,
	"f p r t":   flags.RightPenaltyTopFlag,
	"f p r b":   flags.RightPenaltyBotFlag,
	"f l t":     flags.LeftTopFlag,
	"f l b":     flags.LeftBotFlag,
	"f r t":     flags.RightTopFlag,
	"f r b":     flags.RightBotFlag,
	"f t l 50":  flags.TopLeft50Flag,
	"f t l 40":  flags.TopLeft40Flag,
	"f t l 30":  flags.TopLeft30Flag,
	"f t l 20":  flags.TopLeft20Flag,
	"f t l 10":  flags.TopLeft10Flag,
	"f t 0":     flags.Top0Flag,
	"f t r 10":  flags.TopRight10Flag,
	"f t r 20":  flags.TopRight20Flag,
	"f t r 30":  flags.TopRight30Flag,
	"f t r 40":  flags.TopRight40Flag,
	"f t r 50":  flags.TopRight50Flag,
	"f b l 50":  flags.BotLeft50Flag,
	"f b l 40":  flags.BotLeft40Flag,
	"f b l 30":  flags.BotLeft30Flag,
	"f b l 20":  flags.BotLeft20Flag,
	"f b l 10":  flags.BotLeft10Flag,
	"f b 0":     flags.Bot0Flag,
	"f b r 10":  flags.BotRight10Flag,
	"f b r 20":  flags.BotRight20Flag,
	"f b r 30":  flags.BotRight30Flag,
	"f b r 40":  flags.BotRight40Flag,
	"f b r 50":  flags.BotRight50Flag,
	"f l t 30":  flags.LeftTop30Flag,
	"f l t 20":  flags.LeftTop20Flag,
	"f l t 10":  flags.LeftTop10Flag,
	"f l 0":     flags.Left0Flag,
	"f l b 10":  flags.LeftBot10Flag,
	"f l b 20":  flags.LeftBot20Flag,
	"f l b 30":  flags.LeftBot30Flag,
	"f r t 30":  flags.RightTop30Flag,
	"f r t 20":  flags.RightTop20Flag,
	"f r t 10":  flags.RightTop10Flag,
	"f r 0":     flags.Right0Flag,
	"f r bg 10": flags.RightBot10Flag,
	"f r bg 20": flags.RightBot20Flag,
	"f r bg 30": flags.RightBot30Flag,
}
