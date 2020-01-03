package main

func main() {
	add := func(i int, j int) int {
		return i + j
	}
	r1 := calc(add, 10, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)
}
func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

type calculator func(int, int) int

func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

// 함수 사용은 func 로 사용 기본적으로  func 이름 (변수 변수) 리턴 타입{} 이런식이며,
// 변수 지정은 반대로 r int , result := 2 이런식이다 .
