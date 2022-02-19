package main

import (
	"fmt"
	"time"
)

func main() {
	//Задание 3
	fmt.Println(someConst)
	// taskOne()
	// taskTwo()
}

//Задание 1
func taskOne() {
	for i := 0; i <= 5; i++ {
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
