package main

import "fmt"

func main() {
	/* スライスの初期化とリテラル
	 * スライスのリテラルとは、「長さのない配列リテラル」のようなもの
	 * (配列のリテラル自体は固定長)
	 */

	/* 6個要素を持ったint配列作成し、その参照に対してリテラルを作るようなイメージ */
	prms := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(prms)

	/* 6個の要素を持ったbool配列を作成、その参(以下略) */
	bls := []bool{true, false, true, true, false, true}
	fmt.Println(bls)

	/* structの場合も同じで、構造としては6個の要素を持ったstructの配列リテラルみたいになる */
	strc := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(strc)

	/* スライスを新しく作る際、範囲を省略すると、下限は0、上限は対象のスライス(または配列)の長さが設定される
	 * var a [10]int に対するスライスを作るとき、以下の4つはすべて同じ内容になる
	 * - a[0:10]
	 * - a[:10]
	 * - a[0:]
	 * - a[:]
	 */

	/* prms のスライスを作るとき、下限の値を省略すると0が初期値になるので、
	 * prms[:3] の内容は[2, 3, 5]
	 */
	fmt.Println("prms[:3] -> ", prms[:3])

	/* 上限を省略すると、スライスの長さ(要は末尾)が設定されるので、
	 * prms[4:]の内容は [11, 13] になる
	 */
	fmt.Println("prms[:3] -> ", prms[4:])
}
