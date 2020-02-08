// conversation agent
package bob

import (
	"regexp"
	"strings"
)

// Hey mimics replies from a kid with limited vocabulary
//  all uppercase              ==> "Whoa, chill out!"
//  last character = ?         ==> "Sure."
//  all uppercase and [-1] = ? ==> "Calm down, I know what I'm doing!"
// else whatever

func Hey(remark string) string {
	alpha := regexp.MustCompile(`[a-zA-Z]+`)
	question := regexp.MustCompile(`\?$`)

	newremark := ""
	response := "Whatever."
	isAlpha := alpha.Match([]byte(remark))
	matches := alpha.FindAll([]byte(remark), -1)
	isQuestion := question.Match([]byte(strings.TrimSpace(remark)))

	for _, str := range matches {
		newremark += string(str)
	}
	// response selection "state machine"
	if strings.TrimSpace(remark) == "" {
		return "Fine. Be that way!"
	}

	if isQuestion {
		response = "Sure."
	}

	if isAlpha && strings.ToUpper(newremark) == newremark {
		if !isQuestion {
			response = "Whoa, chill out!"
		} else {
			response = "Calm down, I know what I'm doing!"
		}
	}

	return response
}
