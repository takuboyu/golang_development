package main

import "fmt"

func main() {
	//intのスライスで、長さを3, キャパシティを5
	//makeで初期化する時に長さを3, キャパシティを5の[0, 0, 0]makeを作成
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	//キャパシティの引数を省略すると長さと同じ設定になる
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	//長さが0のスライスを生成

	//二通りのやり方がある
	//0のスライスをメモリに確保
	b := make([]int, 0)
	//nilでメモリに確保しない
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	/*
		簡単な演習
		二つのスライスの違いを体感してみる
	*/
	c = make([]int, 0, 5)
	//c = make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
