package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
LoggingSettings log output
logをファイルに書き込む関数、
基本的なログの設定は以下。
*/
func LoggingSettings(logFile string) {
	//logFileを渡してどういったモードで開くかを設定
	///第二引数：読み書き、なければ作成、追記モード, 第三引数: パーミッション設定
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//画面出力とログファイルに書き込み
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//ログのフラグ設定
	///日付、時間、ファイル名を書き込み
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

/*
logging: ログの処理
他言語のようにライブラリーでinfoなどはなく
最低限のものになっているので必要であればサードパーティの使用検討する必要がある
*/
func main() {
	LoggingSettings("./lesson4/logger/test.log")

	/*
		以下のようなファイルOpenのエラー処理など使うことがある
	*/
	_, err := os.Open("gesge")
	if err != nil {
		log.Fatalln("Exit!!", err)
	}

	//Printはdatetime 文字列を出力される
	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	//Fatalはそこで処理が終了(Exitされる)
	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!!")

	fmt.Println("ok?")
}
