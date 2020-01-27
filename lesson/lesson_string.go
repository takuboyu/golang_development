package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	//stringでcastしてやることでアスキーコードではなく文字列出力できる
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"

	/*
		strings.ReplaceでH->Xに置き換えて文字列をコピーして表示しているので
		後続でサイド出力すると元に戻る点に注意
	*/
	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(s)

	/*
		strings.Containsで文字列に指定文字が含まれているかを調べられる
		true: 含まれている  false: 含まれていない
	*/
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println("Test\n" +
		"Test")
	fmt.Println(`World
	            Test
Test`)

	// " ダブルクウォートなど特殊な文字を表示したいときは前にバックスラッシュを入れてあげる
	fmt.Println("\"")
	// `` バッククウォートで囲むやり方もある
	fmt.Println(`"`)
}
