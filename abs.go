package abs

func SaneAbs (num int) int {
  if num > 0 {
    return num
  } else {
    return -num
  }
}

func InSaneAbs (num int) int {
  y := num>>7
  return (num ^ y) - y
}

func InSaneAbs2 (num int) int {
  y := num>>7
  return (num+y) ^ y
}
func InSaneAbs3 (num int) int {
  return num - (2*num & (num>>7))
}

func InSaneAbs4 (num int) int {
  return ((num >> 7) | 1) * num
}
/*
func SaneWay(b *testing.B) {
  for i := 0; i < b.N; i++ {
    for i, number := range numbers {
	SaneAbs(number)
    }
  }
}
func FastWay1(b *testing.B) {
  for i := 0; i < b.N; i++ {
    for _, number := range numbers {
      InSaneAbs(number)
    }
  }
}

func FastWay2(b *testing.B) {
  for i := 0; i < b.N; i++ {
    for _, number := range numbers {
      InSaneAbs2(number)
    }
  }
}

func FastWay3(b *testing.B) {
  for i := 0; i < b.N; i++ {
    for _, number := range numbers {
      InSaneAbs3(number)
    }
  }
}

func FastWay4(b *testing.B) {
  for i := 0; i < b.N; i++ {
    for _, number := range numbers {
      InSaneAbs4(number)
    }
  }
}

func main () {
        flag.Parse()
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
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
*/
