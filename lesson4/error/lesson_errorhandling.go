package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./README.md")
	if err != nil {
		log.Fatalln("error!!")
	}
	defer file.Close()
	data := make([]byte, 100)
	//読み込んだ数をcountに代入
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("error!!!")
	}
	fmt.Println(count, string(data))
}
