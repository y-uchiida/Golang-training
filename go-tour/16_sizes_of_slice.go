package main

import "fmt"

func main() {
	/* スライス s のサイズ
	 * スライスは、データが入っている数である長さ(length)と、
	 * そのスライスが保持することができるデータの最大数(capacity)という大きさの概念を持っている
	 * length とcapacityは、それぞれ len() と cap() を使用して取得できる */

	/* データ数(length)6, 最大格納数(capacity)6 のスライスを生成 */
	prms := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("primes ... length = %d, capacity = %d \n", len(prms), cap(prms))

	/* [0:0]にして、長さ0の部分集合スライスを切り出す */
	len0 := prms[0:0]
	fmt.Printf("len0 ... length = %d, capacity = %d \n", len(len0), cap(len0))

	/* len0も、もともとのスライスの部分集合としてデータを持っているので、長さを伸ばせばデータを取り出せる */
	middle2 := len0[2:4]
	fmt.Printf("middle2 ... length = %d, capacity = %d \n", len(middle2), cap(middle2))
	fmt.Printf("middle2[0] -> %d, len[1] -> %d\n", middle2[0], middle2[1]) /* prms[2], prms[3] と同じになる */
}
