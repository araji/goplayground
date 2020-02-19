package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type status struct {
	name   string
	played int
	won    int
	draw   int
	loss   int
	points int
}

//Tally reades scores and generates standings
func Tally(in io.Reader, out io.Writer) error {
	myerr := errors.New("tally error")
	stdmap := make(map[string]*status)
	teams := make([]*status, 0, 1)
	validResults := [3]string{"win", "draw", "loss"}

	scores, _ := ioutil.ReadAll(in)
	lines := bytes.Split(scores, []byte("\n"))
	for _, line := range lines {
		if strings.HasPrefix(string(line), "#") || len(line) == 0 {
			continue
		}
		result := bytes.Split(line, []byte(";"))
		validScore := false
		if len(result) == 3 {
			team1 := string(result[0])
			team2 := string(result[1])
			score := string(result[2])
			for _, res := range validResults {
				if res == score {
					validScore = true
				}
			}
			if !validScore {
				return myerr
			}
			if _, ok := stdmap[team1]; !ok {
				stdmap[team1] = &status{}
				stdmap[team1].name = team1
				teams = append(teams, stdmap[team1])
			}
			if _, ok := stdmap[team2]; !ok {
				stdmap[team2] = &status{}
				stdmap[team2].name = team2
				teams = append(teams, stdmap[team2])
			}
			stdmap[team1].played++
			stdmap[team2].played++
			switch score {
			case "win":
				stdmap[team1].won++
				stdmap[team2].loss++
				stdmap[team1].points += 3
			case "loss":
				stdmap[team2].won++
				stdmap[team1].loss++
				stdmap[team2].points += 3
			case "draw":
				stdmap[team1].draw++
				stdmap[team2].draw++
				stdmap[team1].points++
				stdmap[team2].points++
			}
		} else {
			return myerr
		}
	}
	// Sort by alphabetic first then points
	sort.SliceStable(teams, func(i, j int) bool { return teams[i].name < teams[j].name })
	sort.SliceStable(teams, func(i, j int) bool { return teams[i].points > teams[j].points })
	//header then data
	fmt.Fprintf(out, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range teams {
		fmt.Fprintf(out, "%-31s|%3d |%3d |%3d |%3d |%3d\n", v.name, v.played, v.won, v.draw, v.loss, v.points)
	}
	return nil
}
