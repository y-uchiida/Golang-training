package main

import "fmt"

/* 構造体
 * C言語の構造体宣言と少し違う
 * type [構造体の型名] struct { メンバ変数と型 }
 */
type Vertex struct {
	X, Y int
	Z    string
}

var (
	v1  = Vertex{1, 2, "fizz"} // 構造体の型名の後ろに、{...} でメンバの初期値値を列挙する
	v2  = Vertex{Y: 0, X: 10}  // "メンバ名: " の記述を用いると、一部のメンバ変数だけ初期化できる
	v3  = Vertex{}             // なにも初期化しないパターン
	ptr = &v3                  // 構造体も、ポインタを利用できる
)

func main() {
	fmt.Println(v1, ptr, v2, v3)
	fmt.Printf("p: %p\n", ptr)
	fmt.Printf("v3: %p\n", &v3)
}
