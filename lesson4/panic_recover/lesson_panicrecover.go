package main

import (
	"fmt"
)

func thirdPartyConnectDB() {
	panic("Unable to connect database")
}

func save() {
	defer func() {
		//panicしたものをリカバリーしてsに代入
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

/*
panic: 強制終了させる
recover: 強制終了させないでリカバリー処理を行う
go自体はあまり推奨しているものではないらしい
error判定をしてしっかりハンドリングすることを推奨しているようです
*/
func main() {
	save()
	fmt.Println("OK?")
}
