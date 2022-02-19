package task

import (
	"time"
)

//Задание 1
func taskOne() []int {
	var res []int
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second)
		go func(iter int) {
			res = append(res, iter)
		}(i)
	}
	return res
}
