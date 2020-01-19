package main

/*
パッケージというgoが提供しているライブラリ群
*/
import (
	"fmt"
	"os/user"
	"time"
)

/*
initializeが一番初めに実行される
*/

func init() {
	fmt.Println("init")
}

func bazz() {
	fmt.Println("Bazz")
}

/*
メインの実行処理
*/
func main() {
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
}
