package main

import "fmt"

/* bool 型で、3つのを変数を宣言
 * var 変数名 型名 の順番
 * ちなみにboolのデフォルト値はfalse
 */
var c, python, java bool

/* int 型で、ftを宣言して初期化
 * var 変数名 型名 = 初期値 の順番
 * これは特殊だ...
 */
var ft int = 42

func main() {
	/* 関数の中でも、当然変数宣言できる */
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(ft)
}
