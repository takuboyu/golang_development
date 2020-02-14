package main

import "fmt"

/*
new: 初期値を入れない状態でもメモリにポインタを確保してくれる
*/
func main() {
	/*
		var p *int = new(int)
		fmt.Println(p)

		//newを使って宣言しないとnilと出力される
		var p2 *int
		fmt.Println(p2)
	*/

	//newとmakeの違い
	///ポインタ(*)を返すものはnewを返さないものはmakeを使う
	s := make([]int, 0)
	fmt.Printf("%T\n", s)
	m := make(map[string]int)
	fmt.Printf("%T\n", m)
	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p *int = new(int)
	fmt.Printf("%T\n", p)
	var st = new(struct{})
	fmt.Printf("%T\n", st)
}
