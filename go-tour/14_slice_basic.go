package main

import "fmt"

/* slice は、配列の部分集合みたいなもの
 * 配列が固定長なのに対して、sliceは可変長に扱うことができる
 * [][型名] で、その型のsliceを宣言できる([]に要素数が入ったら固定長配列！)
 * []int なら、int 型のsliceになる
 * また、sliceや配列の一部から部分的なsliceを作ることができる
 * a[0:3] なら、slice a の0番目の要素から2番目までの要素を含むスライスになる
 * a[start:end] で、start ~ end - 1までを含める
 */
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	/* primes 配列の1番目から3番目までをふくむslice
	 * {3, 5, 7} になる
	 */
	var s []int = primes[1:4]
	fmt.Println("slice s: ", s)

	/* sliceは参照渡しなので、sliceの値を変えると、元の値も変わる */
	s[1] = 0

	fmt.Println("after changing s[1] = 0...")
	fmt.Println("slice s: ", s)     /* s[1] なので、2つ目の要素が変わっている */
	fmt.Println("primes: ", primes) /* s[1] はprimesでの2番要素なので、primes[2]が変わっている */
}
