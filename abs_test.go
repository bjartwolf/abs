package abs
import (
	"testing"
	"fmt"
)

var posNumbers  = [10]int{32,44,42,0,24,2,4,42,24,32}
var numbers  = [10]int{-32,44,42,0,-24,2,4,42,-24,-32}

func TestSaneAbs(t *testing.T) {
    for i, number := range numbers {
	if(SaneAbs(number) != posNumbers[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, SaneAbs(number), posNumbers[i]))
		}
    }
}
