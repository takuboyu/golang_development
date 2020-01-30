package main

import "fmt"

const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

/*
constantはtype指定がない
なので下記例だとvarの時は宣言時点でoverflowするが
constantではエラーにならない
*/
//var big int = 9223372036854775807 + 1
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big - 1)
}
