package multiple_results

import "fmt"

/* 関数は、複数の返り値を返すことができる
 * 関数 swap は、二つのstring を受け取って、それを入れ替えて、2つの文字列を返す
 */
func swap(x, y string) (string, string) {
	return y, x
}

/* swap() に hello と world を渡し、a とb でそれぞれ入れ替わった値を受け取る
 * 複数の返り値を受け取る際に、 := と記述するのが特徴的
 */
func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
