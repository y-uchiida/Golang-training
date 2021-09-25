package import_package

import (
	"fmt"
	"math"
)

func main() {
	/* 大文字で始まっている関数名？は、パッケージからエクスポートされているので、
	 * 外部から参照・利用できる
	 */
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
}
