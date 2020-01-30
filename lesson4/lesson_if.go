package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "ng"
	}
}

func main() {
	//基本的なif文
	//	num := 6
	//	if num%2 == 0 {
	//		fmt.Println("by 2")
	//	} else if num%3 == 0 {
	//		fmt.Println("by 3")
	//	} else {
	//		fmt.Println("else")
	//	}

	//基本的なand or
	///andは&&
	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
	///orは||
	if x == 10 || y == 20 {
		fmt.Println("||")
	}

	//functionを呼び出す例
	result := by2(10)
	if result == "ok" {
		fmt.Println("greatt")
	}
	///一度変数に代入しているので後続でresultを使える
	fmt.Println(result)

	///一行で書くやり方もある
	if result2 := by2(10); result2 == "ok" {
		fmt.Println("great 2")
	}
	///この書き方の場合後続でresult2は使えない
	///後続で使わないのであれば上記書き方もできる
	fmt.Println(result2)
}
