package main

import "fmt"

func one(x int) {
	x = 1
}

func main() {
	var n int = 100
	one(n)
	fmt.Println(n)
	/*
		fmt.Println(n)

		//&を付けると100が格納されているメモリー領域のアドレスを出力
		fmt.Println(&n)

		///*を付けるとポインタ型として扱われる
		///以下だと実体がintのポインタ型で先ほどのnのアドレスを代入している

		var p *int = &n
		fmt.Println(p)

		///変数pのアドレスの中身(実体)を見たい時には*をつけてあげれば出力される
		///pにはnのアドレスが入っていて(&n)、その中身(実体)なので100が出力される
		fmt.Println(*p)
	*/
}
