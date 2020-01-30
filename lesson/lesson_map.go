package main

import "fmt"

/*
map
pythonでいう辞書型と同じ
*/
func main() {
	m := map[string]int{"apple": 100, "banna": 200}
	fmt.Println(m)
	//valueを取り出す場合
	fmt.Println(m["apple"])
	//valueを代入する場合
	m["banna"] = 300
	fmt.Println(m)
	//新しい要素を付け足す場合
	m["new"] = 500
	fmt.Println(m)
	//mapにないものを取り出そうとするとどうなるか
	fmt.Println(m["nothing"])

	//値が入ってるか確かめる場合
	//ok(第二引数)でbool値判定もできる
	v, ok := m["apple"]
	fmt.Println(v, ok)
	//mapにないものを取り出そうとするとどうなるか
	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	//初期化について
	///メモリ上に空のmapを作ってから値を代入
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	///メモリ上にmapを作っていないとエラーになる(宣言しただけになっている)
	/*
		var m3 = map[string]int
		fmt.Println("nil")
		m3["pc"] = 5000
		fmt.Println(m3)
	*/

}
