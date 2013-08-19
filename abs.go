package abs

func SaneAbs (num int) int {
  if num > 0 {
    return num
  } else {
    return -num
  }
}

func InSaneAbs (num int) int {
  y := num>>63
  return (num ^ y) - y
}

func InSaneAbs2 (num int) int {
  y := num>>63
  return (num+y) ^ y
}
func InSaneAbs3 (num int) int {
  return num - (2*num & (num>>63))
}

func InSaneAbs4 (num int) int {
  return ((num >> 63) | 1) * num
}
