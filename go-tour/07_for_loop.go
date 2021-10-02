package main

import (
	"fmt"
)

func main() {
	/*
	 * 基本的な書式 for i := 1; i <= 100; i++ {  }
	 * forの後ろにかっこが必要ない！
	 */
	for i := 0; i < 10; i++ {
		fmt.Println("i of loop_01 = ", i)
	}

	/*
	 * forの初期化処理と後処理の記述を省略でき、条件文だけを記載することもできる
	 * while ループは、for 文で条件だけ書くことで表現される
	 */
	n, count := 1, 0
	for n < 1000 {
		n += n
		count++
		fmt.Println(count, "times of loop_02 ... n =", n)
	}
}
