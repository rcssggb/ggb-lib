package lexer

import (
	"errors"
	"fmt"
	"strings"
)

// TeamNames parses (ok team_names ...
func TeamNames(m string) (lTeam, rTeam string, err error) {
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
			lTeam = trimmedTeam
		}
		if strings.Contains(team, "team r") {
			trimmedTeam := team
			trimmedTeam = strings.TrimPrefix(trimmedTeam, "team r ")
			rTeam = trimmedTeam
		}
	}

	return
}
