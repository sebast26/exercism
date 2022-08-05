package tournament

import (
	"bufio"
	"errors"
	"io"
	"sort"
	"strings"
	"text/template"
)

const templ = `
{{- printf "%-30s" "Team"}} | MP |  W |  D |  L |  P
{{range .Teams -}}
{{printf "%-30s" .Name}} | {{printf "%2d" .MP}} | {{printf "%2d" .Wins}} | {{printf "%2d" .Deals}} | {{printf "%2d" .Losses}} | {{printf "%2d" .Points}}
{{end -}}
`

var (
	resultsTmpl    = template.Must(template.New("resultsTmpl").Parse(templ))
	ErrInvalidLine = errors.New("invalid line")
)

type team struct {
	Name                    string
	MP, Wins, Deals, Losses int
}

func (t *team) win() {
	t.MP += 1
	t.Wins += 1
}

func (t *team) draw() {
	t.MP += 1
	t.Deals += 1
}

func (t *team) loss() {
	t.MP += 1
	t.Losses += 1
}

func (t team) Points() int {
	return 3*t.Wins + t.Deals
}

func Tally(reader io.Reader, writer io.Writer) error {
	teamMap := make(map[string]team)
	logs, err := readLogs(reader)
	if err != nil {
		return err
	}
	for _, tokens := range logs {
		updateTeams(teamMap, tokens[0], tokens[1], tokens[2])
	}
	return writeResults(teamMap, writer)
}

func readLogs(reader io.Reader) ([][]string, error) {
	logs := make([][]string, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		tokens, err := extractTokens(line)
		if err != nil {
			return nil, err
		}
		logs = append(logs, tokens)
	}
	return logs, scanner.Err()
}

func extractTokens(line string) ([]string, error) {
	tokens := strings.Split(line, ";")
	if len(tokens) != 3 {
		return nil, ErrInvalidLine
	}
	if tokens[2] != "draw" && tokens[2] != "win" && tokens[2] != "loss" {
		return nil, ErrInvalidLine
	}
	return tokens, nil
}

func updateTeams(teams map[string]team, t1 string, t2 string, result string) {
	team1 := findTeam(teams, t1)
	team2 := findTeam(teams, t2)

	switch result {
	case "draw":
		team1.draw()
		team2.draw()
	case "win":
		team1.win()
		team2.loss()
	case "loss":
		team1.loss()
		team2.win()
	}

	teams[t1] = team1
	teams[t2] = team2
}

func findTeam(teams map[string]team, name string) team {
	t, ok := teams[name]
	if !ok {
		t = team{Name: name}
	}
	return t
}

func writeResults(teamMap map[string]team, writer io.Writer) error {
	teams := make([]team, 0, len(teamMap))
	for _, t := range teamMap {
		teams = append(teams, t)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].Points() == teams[j].Points() {
			return teams[i].Name < teams[j].Name
		}
		return teams[i].Points() > teams[j].Points()
	})

	return resultsTmpl.Execute(writer, struct {
		Teams []team
	}{Teams: teams})
}
