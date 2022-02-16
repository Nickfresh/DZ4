package main

import "fmt"

//Функция "Поиск отсутствующего элемента"
func Solution(A []int) int {
	N := len(A)
	switch N {
	case 0:
		return 0 //Массив пустой
	default:
		var m map[int]bool = make(map[int]bool)
		for _, val := range A {
			if m[val] {
				return 0 //Число в массиве встречается более 1 раза
			} else {
				m[val] = !m[val]
			}
		}
		for i := 1; i <= N; i++ {
			if !m[i] {
				return i
			}
		}
		return 0
	}
}

func main() {
	A := []int{2, 3, 1, 5}
	fmt.Println(Solution(A))
}
