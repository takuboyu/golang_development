package main

import "fmt"

//バイト型について
func main() {
	//ASCIIコードを表示
	b := []byte{72, 73}
	fmt.Println(b)
	//stringでcastすれば文字で表示できる
	fmt.Println(string(b))

	//文字列を入れればASCIIコードで返却
	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}
