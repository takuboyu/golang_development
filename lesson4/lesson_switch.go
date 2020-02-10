package main

import "fmt"

func getOsName() string{
	return "aaaa"
}

/*
case文ぽい
*/
func main() {
	os := getOsName()
	switch os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("default!!")
	}

	/*
	変数osをswitch文以降で使わないという事であれば
	一文で書くこともできる
	*/
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!!")
	case "windows":
		fmt.Println("Windows!!")
	default:
		fmt.Println("Mac!!", os)
	}
	//ここではもうosは使えなくなる
	fmt.Println(os)
}
