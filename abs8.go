package abs

func SaneAbs8 (num int8) int8 {
  if num > 0 {
    return num
  } else {
    return -num
  }
}

func InSaneAbs8 (num int8) int8 {
  y := num>>7
  return (num ^ y) - y
}

func InSaneAbs82 (num int8) int8 {
  y := num>>7
  return (num+y) ^ y
}
func InSaneAbs83 (num int8) int8 {
  return num - (2*num & (num>>7))
}

func InSaneAbs84 (num int8) int8 {
  return ((num >> 7) | 1) * num
}
