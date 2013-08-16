package main
import ("fmt"
	"testing"
	"flag"
	"os"
	"log"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func SaneAbs (num int8) int8 {
  if num > 0 {
    return num
  } else {
    return -num
  }
}

func InSaneAbs (num int8) int8 {
  y := num>>7
  return (num ^ y) - y
}

func InSaneAbs2 (num int8) int8 {
  y := num>>7
  return (num+y) ^ y
}
func InSaneAbs3 (num int8) int8 {
  return num - (2*num & (num>>7))
}

func InSaneAbs4 (num int8) int8 {
  return ((num >> 7) | 1) * num
}


func SaneWay(b *testing.B) {
  var a int8
  var c int8
  a = -42
  c = 13
  x := a
	for i := 0; i < b.N; i++ {
    x = a^c^x
    SaneAbs(x)
	}
}
func FastWay1(b *testing.B) {
  var a int8
  var c int8
  a = -42
  c = 13
  x := a
	for i := 0; i < b.N; i++ {
    x = a^c^x
    InSaneAbs(x)
	}
}

func FastWay2(b *testing.B) {
  var a int8
  var c int8
  a = -42
  c = 13
  x := a
	for i := 0; i < b.N; i++ {
    x = a^c^x
    InSaneAbs2(x)
	}
}

func FastWay3(b *testing.B) {
  var a int8
  var c int8
  a = -42
  c = 13
  x := a
	for i := 0; i < b.N; i++ {
    x = a^c^x
    InSaneAbs3(x)
	}
}

func FastWay4(b *testing.B) {
  var a int8
  var c int8
  a = -42
  c = 13
  x := a
	for i := 0; i < b.N; i++ {
    x = a^c^x
    InSaneAbs4(x)
	}
}


func main () {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	fmt.Println("Sane way")
  fmt.Println(testing.Benchmark(SaneWay))
	fmt.Println("Fast way I")
	fmt.Println(testing.Benchmark(FastWay1))
	fmt.Println("Fast way II")
	fmt.Println(testing.Benchmark(FastWay2))
	fmt.Println("Fast way III")
	fmt.Println(testing.Benchmark(FastWay3))
	fmt.Println("Fast way IV")
	fmt.Println(testing.Benchmark(FastWay4))
}
