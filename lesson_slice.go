package main

import "fmt"

func main() {
	//
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4]) //[3 4]
	fmt.Println(n[:2])  //[3 4]
	fmt.Println(n[2:])  //[3 4]
	fmt.Println(n[:])   //[3 4]

	//[][]を二つ並べることでスライスの中にスライスを入れる入れ子のようにできる
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

	/*
		配列とは違いスライスではappendが可能
	*/
	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)
}
