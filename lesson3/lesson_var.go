package main

import "fmt"

/*
varで変数宣言
var宣言した場合はmain関数の外でも宣言可能
*/
func main() {
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)
	/*
		省略宣言することも可能
		この場合は関数内で宣言しないとエラーになる
	*/
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
}
