package main

import (
	"fmt"
	"time"
)

const someConst = "program run on linux"

func main() {
	fmt.Println(someConst)
}

//Задание 1
func taskOne() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		go func(iter int) {
			fmt.Println(iter)
		}(i)
	}
}

//Задание 2
func taskTwo() {
	for i := 0; i < 5; i++ {
		// начала изменения
		i = 5
		fmt.Println("any string")
		// конец изменения
	}
}
