package main

import (
	"fmt"
)

/*
 * 関数内での変数の暗黙的宣言: ":=" を使う
 * 関数の外側では、var ... を利用する
 */
func main() {

	/* := を使ってシンプルにかける
	 * 複数並べて初期化することも可能
	 * 型は、代入した内容から推測される
	 */
	cat, i, j, py := "Mourinette", 4, 2, 3.14

	fmt.Println(i, j, cat, py)
}
