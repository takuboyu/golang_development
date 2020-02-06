package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	//rangeだと上記処理をもっと簡単に書ける
	for i, v := range l {
		fmt.Println(i, v)
	}
	///index番号を使いたくない時は _ で省略もできる
	for _, v := range l {
		fmt.Println(v)
	}

	//mapの値を取り出す時
	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//keyだけ取り出したい時
	for k := range m {
		fmt.Println(k)
	}

	//valueだけ取り出したい時
	for _, v := range m {
		fmt.Println(v)
	}
}
