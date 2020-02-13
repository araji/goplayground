package raindrops
import(
	"fmt"
)
/**
Convert 
- has 3 as a factor, add 'Pling' to the result.
- has 5 as a factor, add 'Plang' to the result.
- has 7 as a factor, add 'Plong' to the result.
- _does not_ have any of 3, 5, or 7 as a factor, the result should be the digits of the number.

*/

func Convert(i int ) string {
	rain := ""
	
	factored := false

	if (i % 3  == 0 ) {
		rain += "Pling"
		factored = true
	} 
	if (i % 5== 0 ) {
		rain += "Plang"
		factored = true
	} 
	if (i % 7 == 0 ) {
		rain+="Plong"
		factored = true
	} 
	
	if !factored {
		return fmt.Sprintf("%d", i)
	} else {
		return rain
	}
}