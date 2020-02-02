package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	//このような書き方もある
	sum := 1
	//セミコロン省略してもOK
	//for sum < 10 {
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
}
