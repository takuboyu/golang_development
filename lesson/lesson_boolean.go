package main

import "fmt"

func main() {
	// var t, f bool = true, false
	t, f := true, false

	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", f, f)
	/*
		あえて強制でbool型でないと表示しないとする場合は%tを追加
		もし他の型を入れるとintですよみたいな感じで教えてくれる
		%vはvalueなので他の値でも入れられる
	*/
	fmt.Printf("%T %v %t\n", t, t, 1)
	fmt.Printf("%T %v %t\n", f, 1, f)

	//論理演算子
	//and
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true && false)

	//or
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(true || false)

	//not
	fmt.Println(!true)
	fmt.Println(!false)

}
