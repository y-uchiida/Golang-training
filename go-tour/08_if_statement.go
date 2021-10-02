package main

import (
	"fmt"
)

func main() {
	/* 基本書式 if [条件] {}
	 * if のあと、条件式は()で囲まない
	 */
	num := 100
	if num > 42 {
		fmt.Println(num, "is larger than 42")
	}

	/* if statement で宣言した変数は、そのifと連携したeleseのブロックのスコープでのみ有効になる */
	if num := -100 - 100; num < 42 { /* numをif statementで再定義 */
		fmt.Println(num, "is smaller than 42")
	}

	fmt.Println("after if statement, num:", num)
}
