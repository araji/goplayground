//twofer 
package twofer
import (
	"fmt"
)
// ShareWith
/**
|Name    |String to return 
|:-------|:------------------
|Alice   |One for Alice, one for me. 
|Bob     |One for Bob, one for me.
|        |One for you, one for me.
|Zaphod  |One for Zaphod, one for me.
*/
func ShareWith(name string) string {
	if len(name) == 0 { 
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
