package main

import "fmt"

func main() {

	fmt.Println("main start")

	for i := 0; i < 10; i++ {
		/* defer statement で処理を遅延実行させる
		 * 式の判定などはその時点で行われるが、実際に処理されるのは
		 * 呼び出し元の関数の処理が完了してから
		 * 複数のdeferが実行された場合、スタック形式で保持される
		 * 最後に先入れ後出し(LIFO)の順で処理する
		 */
		defer fmt.Println(i)
	}

	fmt.Println("main end")
}
