package main

import (
	"fmt"
)

/* ニュートン法で平方根を求める */
func Sqrt(x float64) float64 {
	z := 1.0
	num := 1
	ret := float64(0)
	for {
		ret = (z*z - x) / (2 * z)
		/* 誤差が0.000000000001の範囲になったら終了 */
		if -0.000000000001 < (ret) && (ret) < 0.000000000001 {
			return z
		}
		z -= ret
		fmt.Println("loop", num, ": z =", z)
		num++
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
