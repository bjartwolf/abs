package abs
import (
	"testing"
	"fmt"
	"math"
)

var posNumbers  = [10]int{340,4253656356263,454265436245435,0,24,2324542544354352,4,42,24,math.MaxInt64}
var numbers  = [10]int{-340,-4253656356263,-454265436245435,0,-24,2324542544354352,-4,-42,-24,-math.MaxInt64}

func TestSaneAbs(t *testing.T) {
    for i, number := range numbers {
	if(SaneAbs(number) != posNumbers[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, SaneAbs(number), posNumbers[i]))
		}
    }
}

func TestInSaneAbs(t *testing.T) {
    for i, number := range numbers {
	if(InSaneAbs(number) != posNumbers[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, SaneAbs(number), posNumbers[i]))
		}
    }
}

func TestInSaneAbs2(t *testing.T) {
    for i, number := range numbers {
	if(InSaneAbs2(number) != posNumbers[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, SaneAbs(number), posNumbers[i]))
		}
    }
}

func TestInSaneAbs3(t *testing.T) {
    for i, number := range numbers {
	if(InSaneAbs3(number) != posNumbers[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, SaneAbs(number), posNumbers[i]))
		}
    }
}

func TestInSaneAbs4(t *testing.T) {
    for i, number := range numbers {
	if(InSaneAbs4(number) != posNumbers[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, SaneAbs(number), posNumbers[i]))
		}
    }
}

func BenchmarkSaneAbs(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers {
			SaneAbs(number)
		}
	}
}

func BenchmarkInSaneAbs(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers {
			InSaneAbs(number)
		}
	}
}

func BenchmarkInSaneAbs2(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers {
			InSaneAbs2(number)
		}
	}
}

func BenchmarkInSaneAbs3(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers {
			InSaneAbs3(number)
		}
	}
}

func BenchmarkInSaneAbs4(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers {
			InSaneAbs4(number)
		}
	}
}
