package main

import "fmt"

//無名関数-NoNameFunction-

func main() {
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	func(x, y int) int {
		return x + y
	}(1, 2)
	//関数の最後に(1,2)を追加することで引数を指定できる

}
