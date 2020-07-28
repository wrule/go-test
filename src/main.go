// 猜数字游戏
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guess(dst int) {
	for {
		fmt.Print("请猜1000以内自然数: ")
		var num int
		fmt.Scanf("%d", &num)
		fmt.Println(num)
		if num > dst {
			fmt.Println("大了")
		} else if num < dst {
			fmt.Println("小了")
		} else {
			fmt.Println("猜对了!")
			break
		}
	}
}

func game() {
	rand.Seed(time.Now().UnixNano())
	var dst int
	for {
		dst = rand.Intn(1000)
		fmt.Println("我已经想好了一个数字,你猜吧")
		guess(dst)
	}
}

func main() {
	game()
}
