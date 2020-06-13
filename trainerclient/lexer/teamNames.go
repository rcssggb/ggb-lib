package lexer

import (
	"errors"
	"fmt"
	"strings"
)

type TeamPositions map[string]string

// TeamNames parses (ok team_names ...
func TeamNames(m string) (tp TeamPositions, err error) {
	tp = make(TeamPositions)

	trimmedMsg := m
	trimmedMsg = strings.TrimPrefix(trimmedMsg, "(ok team_names ")
	trimmedMsg = strings.TrimSuffix(trimmedMsg, ")\x00")

	splitTeams := strings.Split(trimmedMsg, ") (")
	if len(splitTeams) > 2 {
		err = errors.New("something went wrong with team_names parsing")
		fmt.Println(err)
		return
	}

	for _, team := range splitTeams {
		team = strings.TrimPrefix(team, "(")
		team = strings.TrimSuffix(team, ")")

		if strings.Contains(team, "team l") {
			trimmedTeam := team
			trimmedTeam = strings.TrimPrefix(trimmedTeam, "team l ")
			tp["left"] = trimmedTeam
		}
		if strings.Contains(team, "team r") {
			trimmedTeam := team
			trimmedTeam = strings.TrimPrefix(trimmedTeam, "team r ")
			tp["right"] = trimmedTeam
		}
	}

	return
}
