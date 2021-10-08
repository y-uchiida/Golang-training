package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // ポインタpにiの値をセット
	fmt.Println(*p) // *をつけると、ポインタがさしている値を表示できる
	*p = 21         // ポインタpを通じて、iの値を更新する
	fmt.Println(i)

	p = &j
	*p = *p / 37 // ポインタ変数を用いて値の演算もできる
	fmt.Println(j)

	var items = [3]int{1, 2, 3}
	p2 := &items
	fmt.Printf("&items:\t %p\n", &items)
	fmt.Printf("p2:\t %p\n", p2)
	/* C言語のように、ポインタの数値自体を変化させるようなことはできない
	 * (ポインタ演算の機能はない)
	 */
	// fmt.Println(p2 + 1) /* これはエラーになる */

}
