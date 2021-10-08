package main

import "fmt"

/*
 * array による配列データの利用
 * var 変数名 [要素数][型名] という記述が必要
 * 固定長配列
 */
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	/* 宣言と同時に値を入れる(初期化する)場合は、
	 * {} で値を囲んで、,で区切って列挙する
	 */
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
