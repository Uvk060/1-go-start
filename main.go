package main

import "fmt"

func main() {
	a := 'Ю'
	var b int = 5 //полный способ, int можно  не писать
	var c = 6
	d := 7
	d, q := 9, 132 //можно присваивать несколько переменнных
	q, d = d, q
	d, j := q, 1000 //если есть хоть 1 новая переменная, то через присвоение :
	c = 15
	fmt.Println("\n", a, b, c, d, q, j)
	jogging()
	numbers()
	evenNumbers()
	exampleSwitch()
}

func jogging() {
	if sprint, marathon := 2, 232; sprint > marathon { //блок if
		println("With mistake")
	} else {
		println("YES, a sprint is shorter than a marathon")
	}
}

func numbers() { //циклы loops
	i := 0
	for i < 10 {
		println(i)
		i++
	}
}
func evenNumbers() { //четные числа
	i := 0
	for {
		if i%2 == 1 {
			i++
			continue
		}
		println(i)
		i++
		if i > 14 {
			break
		}
	}
}

func exampleSwitch() {
	s := 15
	switch s {
	case 10:
		println("ONE")
	case 15:
		println("TWO")
	case 30:
		println("Three")
	default:
		println("Default!!!")

	}
}
