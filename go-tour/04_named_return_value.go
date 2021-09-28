/* コードのコメントは以下から拝借
 * https://www.ushiblo.com/go-09-a-tour-of-go-named-return-values/
 * 要するに、あらかじめ戻り値になる変数を、関数定義と一緒にできるようになってるということかな...
 */

package main

import "fmt"

// (x, y int) が戻り値のパラメータ
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// リターンステートメントの Expression List を省略すると、
	// 戻り値のパラメータに定義された変数が充てがわれるので、こんなイメージ
	// return x, y
	return
}

func main() {
	// ついでにコレも…
	// このサンプルのように Println の引数に split ファンクションを書くと、
	// split ファンクションの戻り値（x と y の値）が Println の引数として扱われるので、こんなイメージ
	// fmt.Println( 7, 10 )
	fmt.Println(split(17))
}
