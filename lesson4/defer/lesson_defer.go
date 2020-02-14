package main

import (
	"fmt"
	"os"
)

/*
defer: 遅延実行
fuctionの実行が終わった後に処理される
*/
func main() {
	/*
		//hello worldの順で出力される
		defer fmt.Println("world")
		fmt.Println("hello")
	*/

	/*
		stacking defer: 初めに入れたものが最後に実行される
		以下のように出力される
		success
		run
		3
		2
		1
	*/
	/*
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/
	/*
		どういう時にdefeerを使うのか
		一例としてファイルを扱う時。
		Openした後にdeferでCloseを処理を行えば
		関数が終わった後にCloseされるのでし忘れを防げる
	*/
	file, _ := os.Open("./README.md")
	defer file.Close()
	data := make([]byte, 100) //Readする為に文字列を扱う際にbyte配列(slice)を作成、ここでは100byte
	file.Read(data)           //Read
	fmt.Println(string(data)) //byteをstringでcastして出力
}
