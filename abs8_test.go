package abs
import (
	"testing"
	"fmt"
)

var posNumbers8  = [10]int8{34,23,42,0,24,2,4,42,24,32}
var numbers8  = [10]int8{-34,-23,-42,0,-24,2,-4,-42,-24,-32}

func TestSaneAbs8(t *testing.T) {
    for i, number := range numbers8 {
	if(SaneAbs8(number) != posNumbers8[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, SaneAbs8(number), posNumbers8[i]))
		}
    }
}

func TestInSaneAbs8(t *testing.T) {
    for i, number := range numbers8 {
	if(InSaneAbs8(number) != posNumbers8[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, InSaneAbs8(number), posNumbers8[i]))
		}
    }
}

func TestInSaneAbs82(t *testing.T) {
    for i, number := range numbers8 {
	if(InSaneAbs82(number) != posNumbers8[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, InSaneAbs82(number), posNumbers8[i]))
		}
    }
}

func TestInSaneAbs83(t *testing.T) {
    for i, number := range numbers8 {
	if(InSaneAbs83(number) != posNumbers8[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, InSaneAbs83(number), posNumbers8[i]))
		}
    }
}

func TestInSaneAbs84(t *testing.T) {
    for i, number := range numbers8 {
	if(InSaneAbs84(number) != posNumbers8[i]) {
		t.Error(fmt.Sprintf("Number is not correct. %d - %d - %d\n", number, InSaneAbs84(number), posNumbers8[i]))
		}
    }
}

func BenchmarkSaneAbs8(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers8 {
			SaneAbs8(number)
		}
	}
}

func BenchmarkInSaneAbs8(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers8 {
			InSaneAbs8(number)
		}
	}
}

func BenchmarkInSaneAbs82(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers8 {
			InSaneAbs82(number)
		}
	}
}

func BenchmarkInSaneAbs83(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers8 {
			InSaneAbs83(number)
		}
	}
}

func BenchmarkInSaneAbs84(b *testing.B) {
	for i:= 0; i< b.N; i++ {
		for _, number := range numbers8 {
			InSaneAbs84(number)
		}
	}
}

