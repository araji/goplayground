package tournament

import (
	"bytes"
	"fmt"
	"errors"
	"strings"
	"io/ioutil"
)

type status struct {
	played int
	won int
	draw int
	loss int
	points int

}


func Tally(in  io.Reader  ,out io.Writer)  error {
	//myerr := errors.New("tally error")
	stdmap :=make(map[string] *status)
	scores,_ :=ioutil.ReadAll(in)
	lines := bytes.Split(scores, []byte("\n"))
	for _, line := range lines {
		result := bytes.Split(line,[]byte(";"))
		if len(result)==3 {
			team1 := string(result[0])
			team2 := string(result[1])
			score := string(result[2])
			
			if _,ok := stdmap[team1] ; !ok {
				stdmap[team1] = &status{}
			}
			if _,ok := stdmap[team2] ; !ok {
				stdmap[team2] = &status{}
			}
			stdmap[team1].played +=1
			stdmap[team2].played +=1
			
			switch score {
			case "win":
				stdmap[team1].won +=1
				stdmap[team2].loss +=1
				stdmap[team1].points +=3
			case "loss":
				stdmap[team2].won +=1
				stdmap[team1].loss +=1
				stdmap[team2].points +=3
			case "draw":
				stdmap[team1].draw +=1
				stdmap[team2].draw +=1
				stdmap[team1].points +=1
				stdmap[team2].points +=1
			}
		}
	}
	fmt.Fprintf(out,"%-32s|%3s |%3s |%3s |%3s |%3s\n","TEAM","MP","W","D","L","P")
	for k,_ := range stdmap {
		fmt.Fprintf(out,"%-32s|%3d |%3d |%3d |%3d |%3d\n",k,stdmap[k].played, stdmap[k].won, stdmap[k].draw, stdmap[k].loss,stdmap[k].points )
	}
	return nil
}

