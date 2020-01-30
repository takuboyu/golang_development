package main

import "fmt"

//functionについて
///func 関数名(第一引数, 第二引数) 返却値の型{}
func add(x int, y int) (int, int) {
	//	fmt.Println("add function")
	//	fmt.Println(x + y)
	return x + y, x - y
}

///返却値に変数名を宣言しておくこともできる
///返却値が複数あった場合に型だけだと可読性にもかけるので変数名を持たせることがある
func cal(price, item int) (result int) {
	result = price * item
	return result
	//返却値の変数名を宣言して入れば変数名は省略可能
	//return
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	//fucn mainの中にfunctionを持たせることもできる
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
	//変数に代入してから呼び出すではなく以下のような書き方もできる
	//並列化処理などで使用することがある
	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}
