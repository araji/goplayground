// acronym  convert provided string to its acronym
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate split and get first letter to Upper
func Abbreviate(s string) string {
	acronym := ""
	//don't forget the + in the re otherwise you would get empty strings in return
	re := regexp.MustCompile(`[^a-zA-Z']+`)
	matches := re.Split(s, -1)
	if len(matches) == 0 {
		return ""
	}
	//TODO: find a better/cleaner way maybe with replaceAllStringFunc
	for _, str := range matches {
		firstChar := string([]rune(str)[0:1])
		acronym += strings.ToUpper(firstChar)
	}
	return acronym
}
