package functions

import "fmt"

/* 関数宣言の書式...
 * func [関数名]([引数名1] [引数1の型], ...) [返り値の型] {}
 * 引数の後ろに型を記述する点が特徴的
 */
func add(x int, y int) int { /* 複数の引数がおなじ型の場合、最後にまとめて型を書くこともできる -> add(x, y int) */
	return x + y
}

func main() {
	fmt.Println(add(42, 42))
}
