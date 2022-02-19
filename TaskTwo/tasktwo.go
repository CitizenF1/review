package task

//Задание 2
func taskTwo() string {
	str := ""
	for i := 0; i < 5; i++ {
		// начала изменения
		i = 5
		str += "any string"
		// конец изменения
	}
	return str
}
