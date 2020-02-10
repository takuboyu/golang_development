package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	//配列を[x]intで型サイズ指定されているので後から追加することはできない(sliceではできる)
	var b [2]int = [2]int{100, 200}
	//配列でappendしようとするとエラーになる
	//b = append(b, 300)
	fmt.Println(b)

}
